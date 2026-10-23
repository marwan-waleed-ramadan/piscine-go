package piscine

func MakeRange(min, max int) []int {
	if min >= max {
		return nil
	}
	answer := make([]int, max-min)

	for i := min; i < max; i++ {
		answer[i-min] = i
	}
	return answer
}
