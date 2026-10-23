package piscine

func TrimAtoi(s string) int {
	num := 0
	isNegative := false
	foundDigit := false

	for _, r := range s {
		if r == '-' && !foundDigit {
			isNegative = true
		}
		if r >= '0' && r <= '9' {
			foundDigit = true
			num = num*10 + int(r-'0')
		}
	}

	if !foundDigit {
		return 0
	}
	if isNegative {
		return -num
	}
	return num
}
