package main

import (
	"github.com/01-edu/z01"
	"os"
)

func BasicAtoi(s string) string {
	x := []rune(s)
	result := os.Args
	//count := 0
	for i := 0; i <= len(x) -1; i++{
		if (x[i]>= 'a') && (x[i]<= 'z') || (x[i]>= 'A') && (x[i]<= 'Z') {
			x[i]= x[i]-13
		}
		result +=string(x[i])
	}
	return result 
func main() {
	result := os.Args
	count := 0
	for range result {
		count++
	}	
	for i := 1; i < count; i++ {
		for _, i2 := range result[i] {
			z01.PrintRune(i2)
		}
		z01.PrintRune(10)
	}
}
