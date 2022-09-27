package piscine

func IsLower(s string) bool {
	r := []rune(s)
	for _, a := range r {
		if a < 'a' || a > 'z' {
			return false
		}
	}
	return true
}
