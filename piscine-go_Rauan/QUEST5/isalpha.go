package piscine

func IsAlpha(s string) bool {
	r := []rune(s)
	for _, a := range r {
		if a < 48 || a > 57 && a < 65 || a > 90 && a < 97 || a > 122 {
			return false
		}
	}
	return true
}
