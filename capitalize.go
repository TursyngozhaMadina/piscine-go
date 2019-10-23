package piscine

func Capitalize(s string) string {
	x := []rune(s)
	resul := ""
	counter := 0
	for range x {
		counter++
	}
	for i := 0; i <= counter-1; i++ {
		if ((x[i] == ' ') || (x[i] == '+')) && (x[i+1] >= 'a') && (x[i+1] <= 'z') {
			x[i+1] = x[i+1] - 32
		}
		resul += string(x[i])
	}
	return resul
}
