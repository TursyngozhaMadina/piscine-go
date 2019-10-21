package piscine

func Sqrt(nb int) int {

	for i := 0; i < 20; i++ {
		if nb == i*i {
			return i
		}
	}
	return 0
}
