package main
import (
	"github.com/01-edu/z01"
	"os"
)
func main () {
	args := os.Args
	count := 0
	//x := []rune()
	for range args {
		count++
	}
	for i := 1; i <= count; i++ {
		if i != 0 {
			x := []rune(args[i])
			for _, r := range x {
				if (x[i]  >= 'A' && x[i]  <= 'Z') || (x[i]  >= 'A' && x[i] <= 'Z') {
					x[i] = x[i] - 32
					z01.PrintRune(r)
				}
				
			}
		}
		z01.PrintRune('\n')
	}

}	
//for _, i2 := range result[i] {