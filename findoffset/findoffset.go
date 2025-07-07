// findoffset.go is a command line tool that finds the offset of the first occurence of a string in a file and prints it to stdout.
package main

import (
	"fmt"
	"math"
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

	filepath, pattern, desired_match_index_as_str := os.Args[1], os.Args[2], os.Args[3]
	target_instance_number, err := strconv.Atoi(desired_match_index_as_str)
	if target_instance_number == 0 {
		fmt.Fprint(os.Stderr, "Can't find the 0th occurence of a string.")
		os.Exit(1)
	}

	// 2. read the file into memory
	// we read the whole file into memory for the sake of the demonstration
	b, err := os.ReadFile(filepath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "read %s: %v", filepath, err)
		os.Exit(1)
	}

	instances := []int{}

	for i := 0; i < len(b)-len(pattern); i++ {
		for j := range pattern {
			if b[i+j] != pattern[j] {
				break
			}
			if j == len(pattern)-1 {
				instances = append(instances, i)
			}
		}
	}
	// The int -> float -> math.Abs -> int cast feels horrific but also whatever.
	if len(instances) == 0 || int(math.Abs(float64(target_instance_number))) > len(instances) {
		fmt.Fprintf(os.Stderr, "Did not find occurence number: %d of %s", target_instance_number, pattern)
		os.Exit(1)
	}
	if target_instance_number > 0 {
		fmt.Fprintf(os.Stdout, "%d\n", instances[target_instance_number-1])
		os.Exit(0)
	}

	if target_instance_number < 0 {
		fmt.Fprintf(os.Stdout, "%d\n", instances[len(instances)+target_instance_number])
		os.Exit(0)
	}

	os.Exit(1)
}
