package piscine

func IsNumeric(str string) bool {
	//x := []rune(str)
	n := true
	for _, cat := range str {
		if !(cat >= '0' && cat <= '9') {
			n = false
		}
	}
	return n
}

//|| cat >= 'A' && cat <= 'Z' || cat >= 'a' && cat <= 'z')
