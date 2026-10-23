package piscine

func BasicAtoi(s string) int {
	result := 0
	for _, ch := range s {
		// Convert the character '0'-'9' to its integer value
		digit := int(ch - '0')
		result = result*10 + digit
	}
	return result
}
