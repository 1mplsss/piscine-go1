package piscine

func IterativeFactorial(nb int) int {
	if nb <= -2147483648 || nb >= 2147483648 || nb < 0 {
		return 0
	} else {
		if nb == 0 || nb == 1 {
			return 1
		} else {
			for i := nb - 1; i >= 1; i-- {
				nb = nb * i
			}
			return nb
		}
	}
}
