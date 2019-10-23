package piscine

func IsAlpha(str string) bool {
	//x := []rune(str)
	n := true
	for _, cat := range str {
		if !( cat >= '0' && cat <= '9' || cat >= 'A' && cat <= 'Z' || cat >= 'a' && cat <= 'z') {
			n = false
		}
	}
	return n
}
