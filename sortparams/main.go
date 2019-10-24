package main

import (
	"github.com/01-edu/z01"
	"os"
)

func main() {
	result := os.Args
	//count := 0
	for j := '!'; j <= '~' ; j++ {
		for i := range result {
            if i != 0  {
				for _, i2 := range result[i] {
					if i2 == j {
						z01.PrintRune(i2)
						z01.PrintRune(10)
					}
				}
			}
		}		  
	}
}
