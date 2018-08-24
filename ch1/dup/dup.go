package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	files := os.Args[1:]
	if len(files) == 0 {
		fmt.Fprintf(os.Stderr, "No input files...\n")
		os.Exit(1)
	} else {
		for _, fileName := range files {
			f, err := os.Open(fileName)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			counts := countLines(f)
			err = f.Close()
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				os.Exit(2)
			}
			for line, n := range counts {
				if n > 1 {
					fmt.Printf("%d\t%s\tFileName: %s\n", n, line, fileName)
				}
			}
		}
	}

}

func countLines(f *os.File) map[string]int {
	counts := make(map[string]int)

	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}

	if input.Err() != nil {
		fmt.Fprintf(os.Stderr, "dup2: %v\n", input.Err())
		os.Exit(2)
	}

	return counts
}
