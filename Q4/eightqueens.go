package main ()

import "github.com/01-edu/z01"

func Check(k int, position int, arr [8]int) bool {
	for i := 0; i < k; i++ {
		test := arr[i]
		if test == position || test == position-(k-i) || test == position+(k-i) {
			return false
		}
	}
	return true
}

func Queens(k int, n int, arr [8]int) {
	if k == n {
		for i := 0; i < n; i++ {
			z01.PrintRune(rune(arr[i] + '1'))
		}
		z01.PrintRune('\n')
	} else {
		for i := 0; i < n; i++ {
			if Check(k, i, arr) {
				arr[k] = i
				Queens(k+1, n, arr)
			}
		}
	}
}

func EightQueens() {
	n := 8
	arr := [8]int{}
	Queens(0, n, arr)
}
