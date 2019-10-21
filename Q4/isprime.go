package piscine

func IsPrime(nb int) bool {
	otvet := true
	if nb > 1 {
		for i := 2; i < nb; i++ {
			if nb%i == 0 {
				otvet = false
			}
		}
	} else {
		otvet = false
	}
	return otvet
}
