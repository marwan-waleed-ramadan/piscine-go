package piscine

func Any(f func(string) bool, a []string) bool {
	for _, x := range a {
		if f(x) {
			return true
		}
	}
	return false
}
