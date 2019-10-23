package piscine

func Join(strs []string, sep string) string {

	var a string
	k := 0
	d := 0
	for ind, let := range strs {
		let = let
		d = ind
	}
	d++
	for let := range strs {
		a += strs[let]
		k++
		if k == d {
			break
		}
		a = a + sep

	}
	return a
}
