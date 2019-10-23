package piscine

func IsAlpha(str string) bool {
	x := []rune(str)
	counter := 0
	n := false
	for range x {
		counter++
	}
	for i := 0; i <= counter-1; i++ {
		if n && (x[i] >= 0 && x[i] <= 47) {
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
