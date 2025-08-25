// findoffset.go is a command line tool that finds the offset of the first occurence of a string in a file and prints it to stdout.
package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	// 1. parse the command line arguments

	// the operating system provides command line arguments to your program.
	// os.Args[0] is the name of the program, and the rest are the 'real' arguments.
	if len(os.Args) != 4 {
		fmt.Fprintf(os.Stderr, "Usage: findoffset <filename> <string> <occurence>")
		os.Exit(1)
	}
	var (
		file        = os.Args[1]
		offset, err = strconv.ParseInt(os.Args[2], 0, 64)
		replacement = os.Args[3]
	)
	if err != nil || offset < 0 {
		fatalf("invalid offset: %v\nUsage: %s <file> <offset> <replacement>", err, os.Args[0])
	}

	f, err := os.OpenFile(file, os.O_RDWR, 0)
	if err != nil {
		fatalf("open %s: %v\n", file, err)
	}
	defer f.Close()

	// copy up until offset
	_, err = io.CopyN(os.Stdout, f, offset)
	if err != nil {
		fatalf("copy: %v\n", err)
	}

	// write replacements to stdout
	_, err = os.Stdout.Write([]byte(replacement))
	if err != nil {
		fatalf("write: %v\n", err)
	}

	// skip replaced bytes
	if _, err := io.CopyN(io.Discard, f, int64(len(replacement))); err != nil {
		fatalf("copy: %v\n", err)
	}

	// copy the rest of the file
	_, err = io.Copy(os.Stdout, f)
	if err != nil {
		fatalf("copy: %v\n", err)
	}
}

func fatalf(format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, format, args...)
	os.Exit(1)
}
