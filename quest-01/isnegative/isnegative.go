package isnegative

import "github.com/01-edu/z01"

//PrintRune IsNegative
func IsNegative(nb int) {

	if nb < 0 {
		
		z01.PrintRune('\')
	} else {
		z01.PrintRune('F')
		z01.PrintRune('\')
	}

	z01.PrintRune(10)

}
