package piscine

func ToLower(s string) string {
	x := []rune(s)
	result := ""
	for i := 0; i <= len(x)-1; i++ {
		if (x[0] >= 'a') && (x[0] <= 'z') {
			x[0] = x[0] + 32
		}
		result += string(x[i])
	}
	}
	return result
}
