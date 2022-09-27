package piscine

func Split(s, sep string) []string {
	seplen := len(sep)
	slen := len(s)
	for i := 0; i < seplen; i++ {
		s += " "
	}
	bool := false
	size := 0
	for i := 0; i < slen; i++ {
		if s[i:i+seplen] == sep && !bool {
			bool = true
			size++
		} else {
			bool = false
		}
	}
	arr := make([]string, size+1)
	word := ""
	j := 0
	for i := 0; i < slen; i++ {
		if s[i:i+seplen] == sep {
			if len(word) == 0 {
				continue
			}
			arr[j] = word
			word = ""
			i = i + seplen - 1
			j++
			continue
		}
		word += string(s[i])
	}
	c := 0
	for i := range word {
		c = i + 1
	}
	if c != 0 {
		arr[j] = word
	}
	return arr
}
