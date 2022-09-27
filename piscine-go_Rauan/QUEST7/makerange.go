package piscine

func MakeRange(min, max int) []int {
	if min >= max {
		return []int(nil)
	}
	rang := max - min
	answer1 := make([]int, rang)

	for i := 0; i < rang; i++ {
		answer1[i] = i + min
	}
	return answer1
}
