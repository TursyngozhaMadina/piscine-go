package piscine

func IsLower(str string) bool {
	//x := []rune(str)
	n := false
	for _, cat := range str {
		if cat >= 'A' && cat <= 'Z' || cat >= 'a' && cat <= 'z' {
			n = true
		}
	}
	return n
}
