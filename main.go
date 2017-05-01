package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
)

func mainError() (err error) {
	http.Handle("/", http.FileServer(http.Dir(".")))
	port := os.Getenv("PORT")
	if port == "" {
		port = "80"
	}
	return http.ListenAndServe(":"+port, nil)
}

func mainCode() int {
	err := mainError()
	if err == nil {
		return 0
	}
	fmt.Fprintf(os.Stderr, "%v: Error: %v\n", filepath.Base(os.Args[0]), err)
	return 1
}

func main() {
	os.Exit(mainCode())
}
