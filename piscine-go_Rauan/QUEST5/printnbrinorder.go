package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	a := []int{}
	if n <= 0 {
		a = append(a, 0)
	}
	for i := 1; n > 0; i++ {
		a = append(a, n%10)
		n /= 10
	}
	for true {
		swapped := false
		for i := 0; i < len(a)-1; i++ {
			if a[i] > a[i+1] {
				swapped = true
				a[i], a[i+1] = a[i+1], a[i]
			}
		}
		if !swapped {
			break
		}
	}
	for i := range a {
		z01.PrintRune(rune('0' + a[i]))
	}
}
