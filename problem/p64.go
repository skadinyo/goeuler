package problem

import (
	"fmt"
	"math"
)

func Problem64() {
	problem64fn(10000)
}

func problem64fn(maxN int) {
	c := 0
	for i := 1; i <= maxN; i++ {
		l := getSqrtPeriodLength(i)
		if 1 == (l % 2) {
			c++
		}
	}
	fmt.Println(c)

}

func getSqrtPeriodLength(x int) int {
	sqrtI := int(math.Floor(math.Sqrt(float64(x))))
	if x == sqrtI*sqrtI {
		return 0
	}
	c := 0

	a := sqrtI
	n := 0
	d := 1

	for a != 2*sqrtI {
		n = d*a - n
		d = (x - (n * n)) / d
		a = (sqrtI + n) / d
		c++
	}

	return c
}
