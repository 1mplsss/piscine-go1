package piscine

func Compare(a, b string) int {
	var c int
	if a == b {
		return 0
	} else if a < b {
		return -1
	} else {
		c = 1
		return c
	}
}
