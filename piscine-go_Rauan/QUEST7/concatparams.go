package piscine

func ConcatParams(args []string) string {
	elem := ""

	for i := 0; i < len(args); i++ {
		elem = elem + args[i]
		if i < len(args)-1 {
			elem = elem + "\n"
		}
	}
	return elem
}
