package piscine

func Index(s string, toFind string) int {
	d := 0
	k := 0
	c := []rune(toFind)
	for i := range toFind {
		if i == i {
			k++
		}
	}
	if k != 0 {
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
	return 0
}
