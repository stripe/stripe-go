// A script that wraps the run of the project test suite and starts stripe-mock
// with a custom OpenAPI + fixtures bundle if one was found in the appropriate
// spot (see `pathSpec` below).
//
// The script passes all its arguments to a Go's test command, and defaults to
// `./...`. For example, both of the following are valid invocations:
//
//	go run test_with_stripe_mock/main.go
//	go run test_with_stripe_mock/main.go ./charge
//
// The reason that we need a separate script for this is because Go's testing
// infrastructure doesn't provide any kind of global hook that we can use to do
// this work in only one place before any tests are run.
package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"time"
)

const (
	defaultStripeMockPort = "12112"

	pathSpec     = "./testing/openapi/spec3.json"
	pathFixtures = "./testing/openapi/fixtures3.json"
)

var portMatch = regexp.MustCompile(` port: (\d+)`)

func main() {
	var stripeMockProcess *os.Process
	autostart := true

	port := os.Getenv("STRIPE_MOCK_PORT")
	if port == "" {
		port = defaultStripeMockPort
	}

	//
	// Maybe start stripe-mock
	//

	autostartEnv := os.Getenv("STRIPE_MOCK_AUTOSTART")
	if autostartEnv == "0" || autostartEnv == "false" {
		fmt.Printf("STRIPE_MOCK_AUTOSTART=%s, "+
			"assuming stripe-mock is already running on port %s\n", autostartEnv, port)
		autostart = false
	} else if _, err := os.Stat(pathSpec); os.IsNotExist(err) {
		fmt.Printf("No custom spec file found, "+
			"assuming stripe-mock is already running on port %s\n", port)
		autostart = false
	}

	if autostart {
		var err error
		port, stripeMockProcess, err = startStripeMock()
		if err != nil {
			exitWithError(err)
		}
	}

	//
	// Run tests
	//

	err := runTests(port)
	if err != nil {
		stopStripeMock(stripeMockProcess)
		exitWithError(err)
	}

	//
	// Stop stripe-mock
	//

	// Try to cleanly stop stripe-mock, but in case we don't, it'll die anyway
	// because it's executing as a subprocess.
	stopStripeMock(stripeMockProcess)
}

//
// Private functions
//

func exitWithError(err error) {
	fmt.Fprintf(os.Stderr, "%v", err)
	os.Exit(1)
}

func runTests(port string) error {
	args := append([]string{"test"}, os.Args[1:]...)

	// Defaults to `./...`, but also allows a specific package (or other CLI
	// flags like `-test.v`) to be passed.
	if len(args) == 1 {
		args = append(args, "./...")
	}

	cmd := exec.Command("go", args...)

	// Inherit this script's environment so that it's still possible to pass
	// the test package flags like `GOCACHE=off`.
	cmd.Env = append(
		os.Environ(),
		"STRIPE_MOCK_PORT="+port,
	)

	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("Error running tests: %v", err)
	}

	return nil
}

func startStripeMock() (string, *os.Process, error) {
	fmt.Printf("Starting stripe-mock...\n")

	cmd := exec.Command(
		"stripe-mock",
		"-https-port", "0", // stripe-mock will select a port
		"-spec", pathSpec,
		"-fixtures", pathFixtures,
	)

	cmd.Stderr = os.Stderr

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return "", nil, fmt.Errorf("Error starting stripe-mock: %v", err)
	}

	err = cmd.Start()
	if err != nil {
		return "", nil, fmt.Errorf("Error starting stripe-mock: %v", err)
	}

	b := make([]byte, 1024)
	var port string

	// We store the entire captured output because the string we're looking for
	// may have appeared across a read boundary.
	var stdoutBuffer bytes.Buffer

	for {
		n, err := stdout.Read(b)
		if err != nil {
			return "", nil, err
		}
		stdoutBuffer.Write(b[0:n])

		// Look for port in "Listening for HTTP on port: 50602"
		matches := portMatch.FindStringSubmatch(stdoutBuffer.String())
		if len(matches) > 0 {
			port = matches[1]
			break
		}

		time.Sleep(100 * time.Millisecond)
	}

	fmt.Printf("Started stripe-mock; PID = %v, port = %v\n", cmd.Process.Pid, port)
	return port, cmd.Process, nil
}

func stopStripeMock(process *os.Process) {
	if process == nil {
		return
	}

	fmt.Printf("Stopping stripe-mock...\n")
	process.Signal(os.Interrupt)
	process.Wait()
	fmt.Printf("Stopped stripe-mock\n")
}
