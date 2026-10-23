package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	if n == -9223372036854775808 {
		z01.PrintRune('-')
		for _, r := range "9223372036854775808" {
			z01.PrintRune(r)
		}
		return
	}
	if n == 0 {
		z01.PrintRune('0')
		return
	}

	if n < 0 {
		z01.PrintRune('-')
		n = -n
	}

	if n >= 10 {
		PrintNbr(n / 10)
	}
	z01.PrintRune(rune('0' + n%10))
}
