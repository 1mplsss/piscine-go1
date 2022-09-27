package piscine

func BasicJoin(elems []string) string {
	s := []rune{}
	for _, elem := range elems {
		a := []rune(elem)
		for _, v := range a {
			s = append(s, v)
		}
	}
	return string(s)
}
