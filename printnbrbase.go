package piscine

import "github.com/01-edu/z01"

func PrintNbrBase(nbr int, base string) {
	if len(base) < 2 {
		z01.PrintRune('N')
		z01.PrintRune('V')
		return
	}

	for i := 0; i < len(base); i++ {
		if base[i] == '+' || base[i] == '-' {
			z01.PrintRune('N')
			z01.PrintRune('V')
			return
		}
		for j := i + 1; j < len(base); j++ {
			if base[i] == base[j] {
				z01.PrintRune('N')
				z01.PrintRune('V')
				return
			}
		}
	}

	var unbr uint64
	if nbr < 0 {
		z01.PrintRune('-')
		unbr = uint64(-(int64(nbr)))
	} else {
		unbr = uint64(nbr)
	}

	baseLen := uint64(len(base))
	if unbr >= baseLen {
		PrintNbrBase(int(unbr/baseLen), base)
	}
	z01.PrintRune(rune(base[unbr%baseLen]))
}
