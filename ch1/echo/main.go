package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	echo1()
	echo2()
}

func echo1() {
	var s, sep string
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

func echo2() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}
