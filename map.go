package piscine

func Map(f func(int) bool, a []int) []bool {
	var ans []bool
	for _, x := range a {
		ans = append(ans, f(x))
	}
	return ans
}
