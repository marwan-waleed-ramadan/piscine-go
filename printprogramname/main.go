package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	name := os.Args[0]
	lastSlash := -1
	for i, r := range name {
		if r == '/' {
			lastSlash = i
		}
	}
	for _, r := range name[lastSlash+1:] {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}
