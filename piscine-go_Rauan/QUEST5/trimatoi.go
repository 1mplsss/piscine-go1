package piscine

func TrimAtoi(s string) int {
	neg := false
	digit := false
	res := 0
	for _, c := range s {
		if c == '-' && !digit {
			neg = true
		}
		if c >= '0' && c <= '9' {
			digit = true
			res = 10*res + int(c-'0')
		}
	}
	if neg {
		return -1 * res
	}
	return res
}
