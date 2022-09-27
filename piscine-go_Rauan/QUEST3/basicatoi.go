package piscine

func BasicAtoi(s string) int {
	toRune := []rune(s)
	var resultArray []int
	for _, v := range toRune {
		switch v {
		case '1':
			resultArray = append(resultArray, 1)
		case '2':
			resultArray = append(resultArray, 2)
		case '3':
			resultArray = append(resultArray, 3)
		case '4':
			resultArray = append(resultArray, 4)
		case '5':
			resultArray = append(resultArray, 5)
		case '6':
			resultArray = append(resultArray, 6)
		case '7':
			resultArray = append(resultArray, 7)
		case '8':
			resultArray = append(resultArray, 8)
		case '9':
			resultArray = append(resultArray, 9)
		default:
			resultArray = append(resultArray, 0)
		}
	}
	return ArrayToInt(resultArray)
}

func ArrayToInt(arrayInt []int) int {
	res := 0
	dec := 1
	for i := len(arrayInt) - 1; i >= 0; i-- {
		res += arrayInt[i] * dec
		dec *= 10
	}
	return res
}
