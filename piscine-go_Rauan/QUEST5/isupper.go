package piscine

func IsUpper(s string) bool {
	r := []rune(s)
	for _, a := range r {
		if a < 'A' || a > 'Z' {
			return false
		}
	}
	return true
}
