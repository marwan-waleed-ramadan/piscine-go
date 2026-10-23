package piscine

func LoafOfBread(str string) string {
	if len(str) == 0 {
		return "\n"
	}
	if len(str) < 5 {
		return "Invalid Output\n"
	}
	count := 0
	skip := false
	result := ""
	for _, char := range str {
		if char == ' ' && !skip {
			continue
		}
		if skip {
			skip = false
			continue
		}
		result += string(char)
		count++
		if count == 5 {
			result += " "
			count = 0
			skip = true
		}
	}
	for i, c := range result {
		if i == len(result)-1 && c == ' ' {
			result = result[:i]
		}
	}
	return result + "\n"
}
