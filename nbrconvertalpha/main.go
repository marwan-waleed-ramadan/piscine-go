package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		return
	}

	upper := args[0] == "--upper"
	if upper {
		args = args[1:]
		if len(args) == 0 {
			return
		}
	}

	for _, a := range args {
		n := 0
		for _, r := range a {
			if r < '0' || r > '9' {
				n = -1
				break
			}
			n = n*10 + int(r-'0')
		}

		if n < 1 || n > 26 {
			z01.PrintRune(' ')
		} else {
			if upper {
				z01.PrintRune(rune(n-1) + 'A')
			} else {
				z01.PrintRune(rune(n-1) + 'a')
			}
		}
	}
	z01.PrintRune('\n')
}
