package piscine

func IsAlpha(str string) bool {
	x := []rune(str)
	re := true
	counter := 0
	for range x {
		counter++
	}
	for i := 0; i <= counter-1; i++ {
		if re && (x[i] >= 'a' && x[i] <= 'z') || (x[i] >= '1' && x[i] <= '9') || (x[i] >= 'A' && x[i] <= 'Z') {
			re = true
		} else {
			re = false
		}
	}
	return re
}
