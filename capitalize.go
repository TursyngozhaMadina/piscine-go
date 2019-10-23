package piscine

func Capitalize(s string) string {
	x := []rune(s)
	resul := ""
	for i := 0; i <= len(x)-1; i++ {
		if ((x[i] == ' ') || (x[i] == '+')) && (x[i+1] >= 'a') && (x[i+1] <= 'z') {
			x[i+1] = x[i+1] - 32
		}
		resul += string(x[i])
	}
	return resul
}
