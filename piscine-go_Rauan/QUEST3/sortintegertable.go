package piscine

func SortIntegerTable(table []int) []int {
	for i := len(table); i > 0; i-- {
		for j := 1; j < i; j++ {
			if table[j-1] > table[j] {
				swapswap := table[j]
				table[j] = table[j-1]
				table[j-1] = swapswap
			}
		}
	}
	return table
}
