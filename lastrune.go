package piscine

func LastRune(s string) rune {
	k := []rune(s)
	count := 0

	for i := range s {
		count = i

	}
	return k[count]
}
