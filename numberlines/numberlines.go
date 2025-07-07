// numberlines reads in a file and prints each line, prepended with an incremented number, to STDOUT
package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: numberlines <filename>")
		os.Exit(1)
	}

	filepath := os.Args[1]
	b, err := os.ReadFile(filepath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "read %s: %v", filepath, err)
		os.Exit(1)
	}

	line_number := 1
	line_start := 0
	for i := 0; i < len(b); i++ {
		if b[i] == '\n' {
			fmt.Fprintf(os.Stdout, "%d. %s\n", line_number, b[line_start+1:i]) // The +1 excludes the newline from getting printed.
			line_start = i
			line_number++
		}
	}
	os.Exit(0)
}
