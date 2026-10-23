package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	summary := make(map[string]int)
	word := ""
	for i := 0; i < len(str); i++ {
		char := str[i]
		if char == ' ' {
			summary[word]++
			word = ""
		} else {
			word += string(char)
		}
	}
	// Add the final word (or empty string if ends with space)
	summary[word]++
	return summary
}
