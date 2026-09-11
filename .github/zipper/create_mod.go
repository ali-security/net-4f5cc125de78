// Command create_mod writes a Go module source zip for a directory using
// golang.org/x/mod/zip.CreateFromDir — Go's own module-zip writer, so the
// member set matches what proxy.golang.org serves by construction.
//
// It lives under .github/ so it is stripped from the tree that gets packed,
// and so it is never part of the published module.
package main

import (
	"log"
	"os"

	"golang.org/x/mod/module"
	"golang.org/x/mod/zip"
)

func main() {
	if len(os.Args) != 5 {
		log.Fatal("usage: create_mod <module-path> <version> <source-dir> <output-zip>")
	}
	m := module.Version{Path: os.Args[1], Version: os.Args[2]}
	f, err := os.Create(os.Args[4])
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	if err := zip.CreateFromDir(f, m, os.Args[3]); err != nil {
		log.Fatal(err)
	}
	log.Printf("created module zip: %s", os.Args[4])
}
