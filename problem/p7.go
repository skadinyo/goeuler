package problem

import (
	"fmt"

	lib "github.com/skadinyo/goeuler/lib"
)

func Problem7() {
	// By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see that the 6th prime is 13.

	// What is the 10 001st prime number?
	problem7fn()
}

func problem7fn() {
	i := 1
	x := 2
	for i != 10001 {
		x++
		if lib.IsPrimeDumb(x) {
			i++
		}
	}
	fmt.Println(x)
}
