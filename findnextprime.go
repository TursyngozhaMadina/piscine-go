package piscine

func FindNextPrime(nb int) int {
	for i := 2; i < nb; i++ {
		if nb%i == 0 {
			return nb + 1
		} else {
			return nb
		}
	}
	return 0
}
