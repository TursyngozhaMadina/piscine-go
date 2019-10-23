package piscine

func IsLower(str string) bool {
	//x := []rune(str)
	n := true
	for _, cat := range str {
		if !(cat >= 'A' && cat <= 'Z' || cat >= 'a' && cat <= 'z') {
			n = false
		}
	}
	return n
}
