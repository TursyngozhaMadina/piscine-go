package main

import (
	"github.com/01-edu/z01"
	"os"
)

func BA(s string) int {
	x := 0
	for _, i := range s {
		if '0' <= i && i <= '9' {
			z := 0
			for k := '1'; k <= i; k++ {
				z++
			}
			x = x*10 + z
		} else {
			x = -1
			break
		}
	}
	if (x != -1) && !(1 <= x && x <= 26) {
		x = -1
	}
	return x
}

func main() {
	arg := os.Args
	pos := 1
	flagupper := false
	c := 0
	for range arg {
		c++
	}
	if c >= 2 && arg[1] == "--upper " {
		pos = 1
		flagupper = true
	}
	for index, k := range arg {
		if index >= pos {
			num := BA(k)
			if num == -1 {
				z01.PrintRune(' ')
			} else {
				if !flagupper {
					z01.PrintRune(rune('a' + num - 1))
				} else {
					z01.PrintRune(rune('A' + num - 1))
				}
			}
		}
	}
	z01.PrintRune(10)
}
