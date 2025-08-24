package Ascii

func GetFileName(banner string) string {
	if banner == "standard" || banner == "standard.txt" {
		return "standard.txt"
	}
	return ""
}
