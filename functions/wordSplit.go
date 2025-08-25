package AsciiArt

func WordSplit(s string) []string {
	slice := []string{}
	str := ""

	for i := 0; i < len(s); i++ {
		if i+1 < len(s) && s[i] == '\\' && s[i+1] == 'n' {
			if len(str) != 0 {
				slice = append(slice, str)
				str = ""
			}
			i = i + 1
			slice = append(slice, "\n")
		} else if s[i] == 10 {
			if len(str) != 0 {
				slice = append(slice, str)
				str = ""
			}
			slice = append(slice, "\n")
		} else {
			str = str + string(s[i])
		}
	}
	if len(str) != 0 {
		slice = append(slice, str)
		str = ""
	}
	return slice
}
