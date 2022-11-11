package problem

import (
	"fmt"
	"math"
)

func Problem9() {
	// A Pythagorean triplet is a set of three natural numbers, a < b < c, for which,

	// a2 + b2 = c2
	// For example, 32 + 42 = 9 + 16 = 25 = 52.

	// There exists exactly one Pythagorean triplet for which a + b + c = 1000.
	// Find the product abc.

	// 1000
	// max c -> 1000/2 -> 500
	//
	problem9fn()
}

func problem9fn() {
	// Brute
	for i := 1; i < 500; i++ {
		for j := 1; j < 500; j++ {
			sr := (math.Sqrt(float64((i * i) + (j * j))))
			srf := math.Floor(sr)
			if sr == math.Sqrt(srf*srf) {
				if i+j+int(srf) == 1000 {
					fmt.Println(i * j * int(srf))
				}
			}
		}
	}
}
