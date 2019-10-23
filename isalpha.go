package piscine

func IsAlpha(str string) bool {
	x := []rune(str)
	re := true
	counter := 0
	for range x {
		counter++
	}
	for i := 0; i <= counter-1; i++ {
		if x[i] == ' ' {
			re = false
		}
	}
	return re
}
