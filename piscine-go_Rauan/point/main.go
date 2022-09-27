package main

import "github.com/01-edu/z01"

type point struct {
	x int
	y int
}

func setPoint(ptr *point, xChosen, yChosen int) {
	ptr.x = xChosen
	ptr.y = yChosen
}

func main() {
	points := &point{}

	setPoint(points, 42, 21)

	string := "x = 42, y = 21"
	for _, i := range string {
		z01.PrintRune(i)
	}
	z01.PrintRune('\n')
}
