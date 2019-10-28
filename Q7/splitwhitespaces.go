package piscine

func SplitWhiteSpaces(str string) []string {
	i := 0
	bykva := true
	for _,word := range str {
		if !(word == ' '  || word == '\n' || word == '\t') {
			bykva = true
		} else {
			if bykva {
				i++
				bykva = false
			}
		}
	}
	a := make([]string, i+1)
	
}
