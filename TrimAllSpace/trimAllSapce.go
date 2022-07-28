package trimallspace

func TrimAllSpace1(s string) string {
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			s = string(append([]byte(s[:i]), []byte(s[i+1:])...))
			i--
		}
	}
	return s
}
func TrimAllSpace2(s string) string {
	var str = []byte(s)
	for i := 0; i < len(str); i++ {
		if str[i] == ' ' {
			str = append(str[:i], str[i+1:]...)
			i--
		}
	}
	return string(str)
}
