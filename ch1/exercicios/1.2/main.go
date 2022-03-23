package main

import (
	"fmt"
	"os"
)

func main() {
	s := ""
	for i := 1; i < len(os.Args); i++ {
		s += fmt.Sprintln(i, os.Args[i])
	}
	fmt.Println(s)
}
