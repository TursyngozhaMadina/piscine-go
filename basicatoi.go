package piscine

func BasicAtoi(s string) int {
	chislo := 0
	for _, cifra := range s {
		desyatok := 0
		for i := '1'; i <= cifra; i++ {
			desyatok++
		}
		chislo = chislo*10 + desyatok
	}
	return chislo
}
