// findoffset.go is a command line tool that finds the offset of the first occurence of a string in a file and prints it to stdout.
package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	var offset, before, after int
	var from string
	var newline bool

	{
		flag.IntVar(&offset, "offset", -1, "offset to read from: must be specified")
		flag.IntVar(&before, "before", 128, "bytes to read before offset: will be clamped to 0")
		flag.IntVar(&after, "after", 128, "bytes to read after offset: will be clamped to 0")
		flag.StringVar(&from, "from", "", "file to read from: if empty, reads from standard input")
		flag.BoolVar(&newline, "newline", false, "append a newline to the output")
		flag.Parse()
	}

	{
		before = max(before, 0)
		before = min(before, offset)
		after = max(after, 0)
		if offset < 0 {
			fmt.Fprintf(os.Stderr, "missing or invalid -offset\n")
			os.Exit(1)
		}
	}

	start := offset - before
	n := before + after
	if n == 0 {
		return
	}
	buf := make([]byte, n)

	f, err := os.Open(from)
	if err != nil {
		fmt.Fprintf(os.Stderr, "open: %s: %v\n", from, err)
		os.Exit(1)
	}
	defer f.Close()

	_, err = f.Seek(int64(start), io.SeekStart)
	if err != nil {
		fmt.Fprintf(os.Stderr, "seek: %s: %v\n", from, err)
		os.Exit(1)
	}

	n, err = io.ReadFull(f, buf)
	if err != nil && err != io.EOF && err != io.ErrUnexpectedEOF {
		fmt.Fprintf(os.Stderr, "read: %s: %v\n", from, err)
		os.Exit(1)
	}

	buf = buf[:n]

	_, err = os.Stdout.Write(buf)
	if err != nil {
		fmt.Fprintf(os.Stderr, "write: %v\n", err)
		f.Close()
		os.Exit(1)
	}
	if newline {
		fmt.Println()
	}
	f.Close()

}
