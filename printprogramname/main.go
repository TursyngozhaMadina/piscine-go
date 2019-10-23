package main

import (
	"github.com/01-edu/z01"
	"os"
)

func main() {

	result := os.Args

	for _, i := range result[0] {

		z01.PrintRune(i)
	}
	z01.PrintRune(10)
}
