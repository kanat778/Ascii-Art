package Ascii

func WordSplit(s string) []string {
	slice := []string{}
	//str := ""

	for i := 0; i < len(s); i++ {
		if i+1 < len(s) && s[i] == '\\' && s[i+1] == 'n' {
			slice = append(slice, "\n")
		}
	}
	return slice
}
