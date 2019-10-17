package piscine

//import "fmt"

func Swap(a *int, b *int) {
	c := *a
	l := *b
	a = l
	b = c 
}
