package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	n := os.Args
	for i, a := range n {
		if i != 0 {
			m := []rune(a)
			for i := range m {
				z01.PrintRune(m[i])
			}
			z01.PrintRune('\n')
		}
	}
}
