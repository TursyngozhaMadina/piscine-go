package main

import "fmt"

func BasicAtoi(s string) int {
	var x int
	for _, y := range s {
		z := 0
		for i := '1'; i < 'y'; i ++ {
			z++
		}
		x = x * 10 + z
	}	
		
return x
}
