// cat reads each file specified on the command ine and writes its contents to standard output.
// usage: cat <file1> [<file2> ...]
package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	for _, file := range os.Args[1:] {
		f, err := os.Open(file)
		if err != nil {
			fmt.Fprintf(os.Stderr, "open %s: %v", file, err)
			os.Exit(1)
		}
		defer f.Close() // io.Copy does this automatically, apparently, but we'll stick with the tutorial and do it the long way
		b, err := io.ReadAll(f)
		if err != nil {
			fmt.Fprintf(os.Stderr, "read %s: %v", file, err)
			os.Exit(1)
		}
		os.Stdout.Write(b)
	}
}
