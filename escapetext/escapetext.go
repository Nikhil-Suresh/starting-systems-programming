// escapetext takes characters with a byte value of <32 or >127 and removes them

// numberlines reads in a file and prints each line, prepended with an incremented number, to STDOUT
package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: escapetext <filename>")
		os.Exit(1)
	}

	filepath := os.Args[1]
	b, err := os.ReadFile(filepath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "read %s: %v", filepath, err)
		os.Exit(1)
	}

	for i := 0; i < len(b); i++ {
		if b[i] < 32 || b[i] > 172 {
			continue
		}
		fmt.Fprintf(os.Stdout, "%s", string(b[i]))
	}
	os.Exit(0)
}
