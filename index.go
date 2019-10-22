package piscine

func Index(s string, toFind string) int {
	d := 0
	//for i := range toFind {
	c := []rune(toFind)
	for ind, book := range s {
		if ind == ind {
			d++
		}
		if c[0] == book && c[0] != 10 {
			return ind
		}
	}
	return -1
}

//	c := []rune(toFind)
//	l := 0
//	n := 0
//	for ind, book := range s {
//		l++
//		if c[0] == book && c[0] != 10 {
//			return ind
