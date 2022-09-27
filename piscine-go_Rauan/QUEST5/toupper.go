package piscine

func ToUpper(s string) string {
	r := []rune(s)
	for a := range r {
		if r[a] >= 'a' && r[a] <= 'z' {
			r[a] = r[a] - 32
		}
	}
	return string(r)
}
