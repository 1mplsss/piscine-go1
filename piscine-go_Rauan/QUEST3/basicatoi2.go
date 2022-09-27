package piscine

func BasicAtoi2(s string) int {
	toRune := []rune(s)
	var resultArray []int
	var one int = 1

	for _, v := range toRune {
		if v != '0' && v != '1' && v != '2' && v != '3' && v != '4' && v != '5' && v != '6' && v != '7' && v != '8' && v != '9' {
			resultArray = append(resultArray, 0)
		} else {
			if v == ' ' {
			}
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
	}
	for _, v := range toRune {
		if v != '0' && v != '1' && v != '2' && v != '3' && v != '4' && v != '5' && v != '6' && v != '7' && v != '8' && v != '9' {
			one *= 0
		} else {
			one *= 1
		}
	}

	if one == 0 {
		return 0
	} else {
		return ArrayToInt2(resultArray)
	}
}

func ArrayToInt2(arrayInt []int) int {
	res := 0
	dec := 1
	for i := len(arrayInt) - 1; i >= 0; i-- {
		res += arrayInt[i] * dec
		dec *= 10
	}
	return res
}
