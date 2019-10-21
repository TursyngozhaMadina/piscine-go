package piscine

func FindNextPrime(nb int) int {

	if nb > 1 {
		if nb%2 == 0 {
			return nb + 1
		} else {
			return nb
		}
	}
	return 0
}
