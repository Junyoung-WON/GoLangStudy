package main

import "fmt"

func main() {
	a := 1
	b := 1

OuterFor:
	for ; a < 10; a++ {
		for b = 1; b < 10; b++ {
			if a*b == 45 {
				break OuterFor
			}
		}
	}
	fmt.Printf("%d * %d = %d\n", a, b, a*b)
}
