package piscine

func StrLen(str string) int {
	nb := []rune(str)
	counter := 0
	for range nb {
		counter++
	}
	return counter
}
