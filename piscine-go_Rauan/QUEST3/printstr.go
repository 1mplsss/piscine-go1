package piscine

import "github.com/01-edu/z01"

func PrintStr(s string) {
	sample := s
	for _, i := range sample {
		z01.PrintRune(i)
	}
}
