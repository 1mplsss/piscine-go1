package piscine

func NRune(s string, n int) rune {
	r := []rune(s)
	if len(s) >= n && n > 0 {
		return r[n-1]
	} else {
		return 0
	}
}
