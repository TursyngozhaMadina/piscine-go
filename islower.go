package piscine

func IsLower(str string) bool {
	//x := []rune(str)
	n := true
	for _, cat := range str {
		if cat >= 33 && cat <= 47 || cat >= 58 && cat <= 64 || cat >= 91 && cat <= 96 || cat >= 123 && cat <= 126 || cat >= 48 && cat <= 57 || cat >= 65 && cat <= 90  {
			n = false
		}
	}
	return n
}
