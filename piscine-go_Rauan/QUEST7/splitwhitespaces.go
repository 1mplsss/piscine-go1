package piscine

func SplitWhiteSpaces(s string) []string {
	rang := 0
	for i, v := range s {
		if v == ' ' && s[i+1] != ' ' {
			rang++
		}
	}
	data := make([]string, rang+1)
	i := 0
	words := ""
	for _, c := range s {
		if c == ' ' || c == '\t' || c == '\n' {
			if words != "" {
				data[i] = words
				words = ""
				i++
			}
		} else {
			words += string(c)
		}
	}
	sz := 0
	for i := range data {
		sz++
		i++
	}
	if words != "" {
		data[sz-1] = words
	}
	return data
}
