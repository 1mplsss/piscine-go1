package piscine

func IsNumeric(s string) bool {
	r := []rune(s)
	for _, a := range r {
		if a < 48 || a > 57 {
			return false
		}
	}
	return true
}
