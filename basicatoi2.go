package piscine

func BasicAtoi2(s string) int {
	result := 0
	for _, ch := range s {
		if ch > '9' || ch < '0' {
			return 0
		} else {
			digit := int(ch - '0')
			result = result*10 + digit
		}
	}
	return result
}
