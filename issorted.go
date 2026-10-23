package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	for i := 0; i < len(a)-2; i++ {
		if f(a[i], a[i+1]) > 0 {
			if f(a[i+1], a[i+2]) < 0 {
				return false
			}
		} else {
			if f(a[i+1], a[i+2]) > 0 {
				return false
			}
		}
	}
	return true
}

func IsGreater(a, b int) int {
	if a > b {
		return 1
	} else if a < b {
		return -1
	}
	return 0
}
