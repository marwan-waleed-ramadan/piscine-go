package piscine

func Atoi(s string) int {
	if s == "" {
		return 0
	}
	result := 0
	sign := 1
	start := 0

	if s[0] == '+' {
		start = 1
	} else if s[0] == '-' {
		sign = -1
		start = 1
	}

	for i := start; i < len(s); i++ {
		ch := s[i]
		if ch < '0' || ch > '9' {
			return 0
		}
		digit := int(ch - '0')
		result = result*10 + digit
	}

	return result * sign
}
