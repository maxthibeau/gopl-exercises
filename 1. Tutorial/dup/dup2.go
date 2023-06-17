package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	countFiles := make(map[string]map[string]bool)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, countFiles)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, countFiles)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			i := 0
			usedFiles := make([]string, len(counts))
			for file := range countFiles[line] {
				usedFiles[i] = file
				i++
			}
			fmt.Printf("%d\t%s\t%s\n", n, line, strings.Join(usedFiles, " "))
		}
	}
}

func countLines(f *os.File, counts map[string]int, countFiles map[string]map[string]bool) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
		if countFiles[input.Text()] == nil {
			countFiles[input.Text()] = map[string]bool{}
		}
		countFiles[input.Text()][f.Name()] = true
	}
	// NOTE: ignoring potential errors from input.Err()
}
