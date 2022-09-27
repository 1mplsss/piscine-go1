package piscine

func ToLower(s string) string {
	r := []rune(s)
	for a := range r {
		if r[a] >= 'A' && r[a] <= 'Z' {
			r[a] = r[a] + 32
		}
	}
	return string(r)
}
