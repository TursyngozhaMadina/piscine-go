package piscine

func IsAlpha(str string) bool {
	x := []rune(str)
	counter := 0
	n := true
	for range x {
		counter++
	}
	for i := 0; i <= counter-1; i++ {
		if n && (x[i] >= 33 && x[i] <= 47) {
			n = false
		} else if n && (x[i] >= 58 && x[i] <= 64) {
			n = false
		} else if x[i] >= 91 && x[i] <= 96 {
			n = false
		} else if x[i] >= 123 && x[i] <= 127 {
			n = false
		} else {
			n = true
		}
	}
	return n
}
