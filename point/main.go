package main

import "github.com/01-edu/z01"

type point struct {
	x int
	y int
}

func setPoint(p *point) {
	p.x = 42
	p.y = 21
}

func main() {
	p := &point{}
	setPoint(p)

	output := []rune{
		120, // x
		32,  // space
		61,  // =
		32,  // space
		52,  // 4
		50,  // 2
		44,  // ,
		32,  // space
		121, // y
		32,  // space
		61,  // =
		32,  // space
		50,  // 2
		49,  // 1
		10,  // newline
	}

	for _, r := range output {
		z01.PrintRune(r)
	}
}
