package piscine

func BasicJoin(strs []string) string {

	var a string
	//	var toConcat[]string
	for b := range strs {
		a += strs[b]
	}
	return a
}
