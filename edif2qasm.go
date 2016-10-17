// edif2qubits is a program that converts an EDIF file into a set of
// qubit weights and coupler strengths.
package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"path"
)

var notify *log.Logger // Help notify the user of warnings and errors.

type Empty struct{} // Zero-byte type for defining and manipulating sets

func main() {
	// Parse the command line.
	var err error
	progName := path.Base(os.Args[0])
	notify = log.New(os.Stderr, progName+": ", 0)
	var r io.Reader
	switch len(os.Args) {
	case 1:
		r = os.Stdin
	case 2:
		f, err := os.Open(os.Args[1])
		if err != nil {
			notify.Fatal(err)
		}
		defer f.Close()
		r = f
	default:
		fmt.Fprintf(os.Stderr, "Usage: %s [<input.edif>]\n", progName)
		os.Exit(1)
	}

	// Parse the specified EDIF file into a top-level s-expression.
	parsed, err := ParseReader(progName, r)
	if err != nil {
		notify.Fatal(err)
	}
	top, ok := parsed.(EdifSExp)
	if !ok {
		notify.Fatalf("Failed to parse the input as an s-expression")
	}
	_ = top // Temporary
}
