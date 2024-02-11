package main

import (
	"fmt"
	"net/http"
	"os"
)

const (
	fileServerDir = "./public"

	servePort = 8080
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "program exited: %s\n", err)
		os.Exit(1)
	}
}

func run() error {
	// Check if the directory exists
	_, err := os.Stat(fileServerDir)
	if os.IsNotExist(err) {
		return fmt.Errorf("unable to serve directory '%s' as it doesn't exist", fileServerDir)
	}

	fileServer := http.FileServer(http.Dir(fileServerDir))

	http.Handle("/", fileServer)

	// Start the server on port 8080
	fmt.Printf("Serving directory '%s' at http://localhost:%d\n", fileServerDir, servePort)
	return http.ListenAndServe(fmt.Sprintf(":%d", servePort), nil)
}
