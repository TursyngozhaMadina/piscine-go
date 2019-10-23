package piscine

func IsPrintable(str string) bool {
	//x := []rune(str)
	n := true
	for _, cat := range str {
		if cat >= 7 && cat <= 13 {
			n = false
		}
	}
	return n
}
