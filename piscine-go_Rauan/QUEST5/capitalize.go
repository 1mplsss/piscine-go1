package piscine

func Capitalize(s string) string {
	a := []rune(s)
	ok := true
	for i, v := range s {
		if ((v >= 'a' && v <= 'z') || (v >= 'A' && v <= 'Z') || (v >= '0' && v <= '9')) && ok {
			if v >= 'a' && v <= 'z' {
				a[i] -= 32
			}
			ok = false
		} else if v >= 'A' && v <= 'Z' {
			a[i] += 32
		} else if !((v >= 'a' && v <= 'z') || (v >= 'A' && v <= 'Z') || (v >= '0' && v <= '9')) {
			ok = true
		}
	}
	return string(a)
}
