package main

import (
	"github.com/01-edu/z01"
	"os"
)

func main() {

	//result := os.Args

	for _, i := range os.Args {
		for _, i2 := range i {
			z01.PrintRune(i2)

		}
		z01.PrintRune(10)

	}
	//z01.PrintRune(10)
}
