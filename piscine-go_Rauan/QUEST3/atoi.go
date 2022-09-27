package piscine

func Atoi(s string) int {
	toRune := []rune(s)
	var resultArray []int
	isNegative := false

	for i, v := range toRune {
		if i == 0 && v == '+' {
			isNegative = false
		} else if i == 0 && v == '-' {
			isNegative = true
		} else if v < '0' || v > '9' {
			resultArray = resultArray[:0]
			resultArray = append(resultArray, 0)
			break
		} else if v == '1' {
			resultArray = append(resultArray, 1)
		} else if v == '2' {
			resultArray = append(resultArray, 2)
		} else if v == '3' {
			resultArray = append(resultArray, 3)
		} else if v == '4' {
			resultArray = append(resultArray, 4)
		} else if v == '5' {
			resultArray = append(resultArray, 5)
		} else if v == '6' {
			resultArray = append(resultArray, 6)
		} else if v == '7' {
			resultArray = append(resultArray, 7)
		} else if v == '8' {
			resultArray = append(resultArray, 8)
		} else if v == '9' {
			resultArray = append(resultArray, 9)
		} else {
			resultArray = append(resultArray, 0)
		}
	}
	if isNegative {
		return -ArrayToIntTest(resultArray)
	} else {
		return ArrayToIntTest(resultArray)
	}
}

func ArrayToIntTest(arrayInt3 []int) int {
	res := 0
	dec := 1
	for i := len(arrayInt3) - 1; i >= 0; i-- {
		res += arrayInt3[i] * dec
		dec *= 10
	}

	return res
}
