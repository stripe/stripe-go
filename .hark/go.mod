// This directory holds hark's changelog data, not Go code.
//
// A go.mod here marks .hark as a nested module, which is the only thing that keeps it
// out of the module zip `go get` downloads: golang.org/x/mod/zip omits "directories
// containing go.mod files", and otherwise drops nothing by name except .bzr, .git,
// .hg and .svn.
//
// Nothing imports this module and it is never published, so the path below only has
// to be syntactically valid and not collide with anything real.
module stripe-go-changelog-data

go 1.24
