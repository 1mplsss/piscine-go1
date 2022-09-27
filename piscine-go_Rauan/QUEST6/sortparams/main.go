package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	name := os.Args[1:]
	for true {
		swapped := false
		for i := 0; i < len(name)-1; i++ {
			if name[i] > name[i+1] {
				name[i], name[i+1] = name[i+1], name[i]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
	for _, elem := range name {

		for _, c := range elem {
			z01.PrintRune(rune(c))
		}
		z01.PrintRune('\n')
	}
}
