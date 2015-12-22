// main.go
// Specification: 	dup2 can read from the standard input or handle a list of
// 					file names, using os.Open to open each one.
package main

import (
	"bufio"
	"fmt"
	"os"
)

type NameCount struct {
	filename string
	counts   map[string]int
}

func NewNameCount() *NameCount {
	var nnc NameCount
	nnc.counts = make(map[string]int)
	return &nnc
}

func main() {
	var ncs []NameCount
	files := os.Args[1:]
	if len(files) == 0 {
		counts := NewNameCount()
		counts.filename = "stdin"
		countLines(os.Stdin, counts)
		ncs = append(ncs, *counts)
	} else {
		for _, file := range files {
			counts := NewNameCount()
			f, err := os.Open(file)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2 error: %v\n", err)
				continue
			}
			countLines(f, counts)
			counts.filename = file
			ncs = append(ncs, *counts)
			f.Close()
		}
	}

	for _, item := range ncs {
		fmt.Printf("Filename: %v\n", item.filename)
		for line, count := range item.counts {
			if count > 1 {
				fmt.Printf("\t%v: %v\n", count, line)
			}
		}
	}
}

func countLines(f *os.File, counts *NameCount) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts.counts[input.Text()]++
	}
}
