// A script that tries to make sure that all API clients (structs called
// `Client`) defined throughout all subpackages are included in the master list
// as a field on the `client.API` type. Adding a new client to `client.API` has
// historically been something that's easily forgotten.
package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
)

func main() {
	//
	// DEBUGGING
	//
	// As you can see, working with Go ASTs is quite verbose, and made not
	// great by all the type casting that's going on. The official docs for the
	// packages provide only the most basic information about how to use them.
	//
	// BY FAR the easiest way to debug something is to just print the AST node
	// you're interested in (it takes in an `interface{}`). The output is
	// verbose, detailed, and extremely informative. For example:
	//
	//     ast.Print(fset, f)
	//

	fset := token.NewFileSet()

	//
	// Step 1: See what clients are in `client.API`
	//

	// Returned as a map to facilitate fast set checking.
	packagesInClientAPI, err := getClientAPIPackages(fset)
	if err != nil {
		exitWithError(err)
	}

	//
	// Step 2: See what clients are in `client.API`
	//

	var packagesWithClient []string
	err = filepath.Walk(".", func(path string, f os.FileInfo, _ error) error {
		if filepath.Ext(path) != ".go" {
			return nil
		}

		if strings.HasSuffix(filepath.Base(path), "_test.go") {
			return nil
		}

		packageName, err := findClientType(fset, path)
		if err != nil {
			exitWithError(err)
		}

		if packageName == nil {
			return nil
		}

		// Prepend a directory so that we know where nested packages are
		// nested.
		//
		// For example, `session` will become `checkout/session`.
		relativeDir := filepath.Dir(path)

		// We're not interested in the immediate parent directory because that
		// has the same name as the package itself, so strip that off the end.
		relativeDir = strings.TrimSuffix(relativeDir, *packageName)

		if relativeDir != "" {
			relativePackageName := relativeDir + *packageName
			packageName = &relativePackageName
		}

		packagesWithClient = append(packagesWithClient, *packageName)
		return nil
	})
	if err != nil {
		exitWithError(err)
	}

	if len(packagesWithClient) < 1 {
		panic("Found no packages with clients; something went wrong " +
			"(maybe check the working directory?)")
	}

	sort.Strings(packagesWithClient)

	var anyMissing bool
	for _, packageName := range packagesWithClient {
		_, ok := packagesInClientAPI[packageName]
		if !ok {
			if !anyMissing {
				fmt.Fprintf(os.Stderr, "!!! the following clients are missing from client.%s in %s\n",
					typeNameAPI, clientAPIPath)
				anyMissing = true
			}

			fmt.Fprintf(os.Stderr, "%s.Client\n", packageName)
		}
	}

	if anyMissing {
		os.Exit(1)
	}
}

//
// Private
//

const (
	// Names of some of the Go types in stripe-go that we're interested in.
	typeNameAPI    = "API"
	typeNameClient = "Client"

	// Path to file containing the `API` type.
	clientAPIPath = "./client/api.go"
)

func exitWithError(err error) {
	fmt.Fprintf(os.Stderr, "%v\n", err)
	os.Exit(1)
}

// findType looks for a Go type defined in the given AST and returns its
// specification.
func findType(f *ast.File, typeName string) *ast.TypeSpec {
	for _, decl := range f.Decls {
		genDecl, ok := decl.(*ast.GenDecl)

		// We're only looking for `Client` structs which are always `GenDecl`
		// of type `TYPE`.
		if !ok || genDecl.Tok != token.TYPE {
			continue
		}

		if len(genDecl.Specs) > 1 {
			panic("Expected only a single ast.Spec for GenDecl with Tok of Type")
		}

		typeSpec, ok := genDecl.Specs[0].(*ast.TypeSpec)
		if !ok {
			panic("Expected ast.TypeSpec in GenDecl with Tok of Type")
		}

		if typeSpec.Name.Name != typeName {
			continue
		}

		// Type names are unique with packages, so return early.
		return typeSpec
	}

	return nil
}

// findClientType looks for a type called `Client` in the `.go` file at the
// target path. If found, it returns the name of the file's defined package.
func findClientType(fset *token.FileSet, path string) (*string, error) {
	source, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	f, err := parser.ParseFile(fset, "", source, 0)
	if err != nil {
		return nil, err
	}

	packageName := f.Name.Name

	typeSpec := findType(f, typeNameClient)
	if typeSpec == nil {
		return nil, nil
	}

	return &packageName, nil

}

// getClientAPIPackages parses the master `API` type found in `client/api.go`,
// looks for all fields that have a type named `Client`, and returns a map
// containing the packages of those clients.
func getClientAPIPackages(fset *token.FileSet) (map[string]struct{}, error) {
	packagesInClientAPI := make(map[string]struct{})

	source, err := ioutil.ReadFile(clientAPIPath)
	if err != nil {
		return nil, err
	}

	f, err := parser.ParseFile(fset, "", source, 0)
	if err != nil {
		return nil, err
	}

	// Regular expression that targets just the last segment(s) of the package
	// path like `coupon` or `issuing/card`. Note that double quotes on either
	// side are also stripped.
	packagePathRE := regexp.MustCompile(`"github.com/stripe/stripe-go/v[0-9]+/(.*)"`)

	// First we need to make a map of any packages that are being imported
	// in ways that don't map perfectly well with the package paths that we'll
	// extract by looking for clients in `.go` files.
	//
	// The first case are packages imported under aliases. We often do this for
	// namespaced packages that have a high probability of collision with
	// something else. For example, `issuing/card` gets imported as
	// `issuingcard`.
	//
	// The second case are nested packages that are referenced are still
	// referenced by their local package name.  For example,
	// `reporting/reportrun` might have been aliased as `reportingreportrun` if
	// things were perfectly consisted, but its package name is already unique
	// enough that no one bothered to do so.
	importAliases := make(map[string]string)
	for _, importSpec := range f.Imports {
		path := importSpec.Path.Value

		// Trim to just the last segment(s) of the package path like `coupon`
		// or `issuing/card`.
		path = packagePathRE.ReplaceAllString(path, "$1")

		// A non-nil `Name` is an alias. Save the alias to our map with the
		// relative package path.
		if importSpec.Name != nil {
			importAliases[importSpec.Name.Name] = path
			continue
		}

		parts := strings.Split(path, "/")

		// A top-level packaage. No need to keep an alias around for it.
		if len(parts) == 1 {
			continue
		}

		// Otherwise, store the alias as the last component of the path keyed
		// to the entire relative path.
		importAliases[parts[len(parts)-1]] = path
	}

	typeSpec := findType(f, typeNameAPI)
	if typeSpec == nil {
		return nil, fmt.Errorf("No 'API' type found in '%s'", clientAPIPath)
	}

	structType, ok := typeSpec.Type.(*ast.StructType)
	if !ok {
		panic(fmt.Sprintf("Expected %s type to be a struct", typeNameAPI))
	}

	for _, field := range structType.Fields.List {
		// A "StarExpr" is a pointer like `*charge.Client`. The type
		// `charge.Client` is wrapped within it.
		//
		// We expect all clients to be pointers, so skip any defined fields
		// that aren't one.
		starExpr, ok := field.Type.(*ast.StarExpr)
		if !ok {
			continue
		}

		selectorExpr, ok := starExpr.X.(*ast.SelectorExpr)
		if !ok {
			continue
		}

		// Only look for fields with types named 'Client'
		if selectorExpr.Sel.Name != typeNameClient {
			continue
		}

		ident, ok := selectorExpr.X.(*ast.Ident)
		if !ok {
			return nil, fmt.Errorf("Expected client field with type '%s' in '%s' "+
				"to be proceeded by an *ast.Ident", typeNameClient, typeNameAPI)
		}

		packageName := ident.Name
		packagesInClientAPI[packageName] = struct{}{}
	}

	for alias, packageName := range importAliases {
		_, ok := packagesInClientAPI[alias]
		if !ok {
			continue
		}

		delete(packagesInClientAPI, alias)
		packagesInClientAPI[packageName] = struct{}{}
	}

	return packagesInClientAPI, nil
}
