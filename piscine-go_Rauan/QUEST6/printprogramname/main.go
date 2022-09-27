package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	name := os.Args[0]
	arr := []rune(name)
	for i := range arr {
		if i > 1 {
			z01.PrintRune(arr[i])
		}
	}
	z01.PrintRune('\n')
}
