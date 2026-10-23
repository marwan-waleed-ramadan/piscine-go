package piscine

func Abort(a, b, c, d, e int) int {
	n1 := a
	n2 := b
	n3 := c
	n4 := d
	n5 := e

	numbers := []int{n1, n2, n3, n4, n5}

	for i := 0; i < 5; i++ {
		for j := 0; j < 4; j++ {
			if numbers[j] > numbers[j+1] {
				temp := numbers[j]
				numbers[j] = numbers[j+1]
				numbers[j+1] = temp
			}
		}
	}

	return numbers[2]
}
