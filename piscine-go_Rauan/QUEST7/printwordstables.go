package piscine

import "github.com/01-edu/z01"

func PrintWordsTables(a []string) {
	for _, d := range a {
		m := []rune(d)
		for i := range m {
			z01.PrintRune(m[i])
		}
		z01.PrintRune('\n')

	}
}
