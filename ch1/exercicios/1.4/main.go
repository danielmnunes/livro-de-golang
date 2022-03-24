package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, "-1")
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2:%v\n", err)
			}
			countLines(f, counts, arg)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func countLines(f *os.File, counts map[string]int, filename string) {
	if filename == "-1" {
		input := bufio.NewScanner(f)
		for input.Scan() {
			counts[input.Text()]++
		}
	} else {
		input := bufio.NewScanner(f)
		for input.Scan() {
			line := fmt.Sprint(filename, ":", input.Text())
			counts[line]++
		}
	}
}
