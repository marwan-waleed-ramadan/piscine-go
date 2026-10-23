package piscine

func Capitalize(s string) string {
	runes := []rune(s)
	inWord := false

	for i := 0; i < len(runes); i++ {
		if (runes[i] >= 'A' && runes[i] <= 'Z') || (runes[i] >= 'a' && runes[i] <= 'z') || (runes[i] >= '0' && runes[i] <= '9') {
			if !inWord {
				if runes[i] >= 'a' && runes[i] <= 'z' {
					runes[i] = runes[i] - ('a' - 'A')
				}
				inWord = true
			} else {
				if runes[i] >= 'A' && runes[i] <= 'Z' {
					runes[i] = runes[i] + ('a' - 'A')
				}
			}
		} else {
			inWord = false
		}
	}

	return string(runes)
}
