package piscine

func LastRune(s string) rune {
	var char rune
	for _, r := range s {
		char = r
	}
	return char
}
