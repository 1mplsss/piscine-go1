package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	name := os.Args[1:]
	offset := 0
	entered := false
	for i := range name {
		if i == 0 && name[i] == "--upper" {
			offset -= 32
			name = name[1:]
		}
	}
	for _, item := range name {
		entered = true
		num := 0
		for _, v := range item {
			num = 10*num + int(v-'0')
		}
		if num >= 1 && num <= 26 {
			z01.PrintRune('a' + rune(num+offset) - 1)
		} else {
			z01.PrintRune(' ')
		}
	}
	if entered {
		z01.PrintRune('\n')
	}
}
