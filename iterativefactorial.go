package piscine

func IterativeFactorial(nb int) int {
	if nb >= 1 {
		result := 1 
		for i := 1 ; i <= nb ; i++ {
			result = result * i
		}
		return result
	} else {

		return 0
	}

}
