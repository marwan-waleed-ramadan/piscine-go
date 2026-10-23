package piscine

func NRune(s string, n int) rune {
	for index, r := range s {
		if n == index+1 {
			return r
		}
	}
	return 0
}
