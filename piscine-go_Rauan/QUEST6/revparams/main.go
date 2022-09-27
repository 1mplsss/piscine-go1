package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	name := os.Args
	for i := len(name) - 1; i >= 1; i-- {
		cur := []rune(name[i])
		for i := 0; i < len(cur); i++ {
			z01.PrintRune(cur[i])
		}
		z01.PrintRune('\n')
	}
}
