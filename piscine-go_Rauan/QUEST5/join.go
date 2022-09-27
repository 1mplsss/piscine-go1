package piscine

func Join(strs []string, sep string) string {
	delim := []rune(sep)
	res := []rune{}
	cnt := 0
	for _, str := range strs {
		cnt++
		a := []rune(str)
		for _, v := range a {
			res = append(res, v)
		}
		if cnt != len(strs) {
			for _, v := range delim {
				res = append(res, v)
			}
		}
	}
	return string(res)
}
