package main

import (
	"github.com/01-edu/z01"
	"os"
)

func main() {
	result := os.Args
	count := 0
	for range result {
		count++
	}
	for i := 1; i < count-1; i++ {
		for _, i2 := range result[i] {
			z01.PrintRune(i2)
		}
		z01.PrintRune(10)
	}
}
