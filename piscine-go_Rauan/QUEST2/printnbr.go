package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	n_start := n
	if n == 0 {
		z01.PrintRune(rune(0 + 48))
		return
	}
	var last int
	lst := []int{}

	for {
		if n == 0 {
			break
		}
		last = n % 10
		n = n / 10
		lst = append(lst, last)
	}

	if n_start > 0 {
		for i := len(lst) - 1; i >= 0; i-- {
			z01.PrintRune(rune(lst[i]) + 48)
		}
	} else {
		z01.PrintRune('-')
		for i := len(lst) - 1; i >= 0; i-- {
			z01.PrintRune(rune(lst[i]*(-1) + 48))
		}
	}
}
