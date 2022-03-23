// Echo1 exibe seus argumentos de linha de comando.
package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	echo1()
	echo2()
	echo3()
	/*
		echo1 -> 8.601µs
		echo2 -> 2.963µs
		echo3 -> 802ns
	*/
}

func echo1() {
	start := time.Now()
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = "\n"
	}
	fmt.Println(time.Since(start))
}

func echo2() {
	start := time.Now()
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = "\n"
	}
	fmt.Println(time.Since(start))
}

func echo3() {
	start := time.Now()
	strings.Join(os.Args[1:], "\n")
	fmt.Println(time.Since(start))
}
