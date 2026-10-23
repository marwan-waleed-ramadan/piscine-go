package piscine

func Split(s, sep string) []string {
	var result []string
	word := ""
	i := 0

	for i < len(s) {
		if len(sep) > 0 && i+len(sep) <= len(s) && s[i:i+len(sep)] == sep {
			result = append(result, word)
			word = ""
			i += len(sep)
		} else {
			word += string(s[i])
			i++
		}
	}

	result = append(result, word)
	return result
}
