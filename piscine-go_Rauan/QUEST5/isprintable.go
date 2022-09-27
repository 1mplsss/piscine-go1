package piscine

func IsPrintable(s string) bool {
	r := []rune(s)
	if len(s) == 0 {
		return true
	}

	for _, a := range r {
		if a >= 0 && a <= 31 {
			return false
		}
	}

	return true
}
