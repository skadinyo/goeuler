package problem

import (
	"fmt"

	lib "github.com/skadinyo/goeuler/lib"
)

func Problem686() {
	problem686fn(123, 678910)
}

func problem686fn(l, n int) int {
	maxLength := len(lib.IntToArray(l))
	maxLengthCalc := maxLength * 8 //dumbass
	d := 1
	drem := 1

	for j := 0; j < maxLength; j++ {
		d *= 10
	}
	for j := 0; j < maxLengthCalc; j++ {
		drem *= 10
	}

	fmt.Println(l, n, maxLength, maxLengthCalc, d, drem)
	x := 1
	c := 0
	counter := 0
	for c != n {
		x *= 2
		//fmt.Println(counter, x, l)
		for x > drem {
			x /= 10
		}
		//get value by maxLengthCalc
		tmp := x
		for tmp > d {
			tmp /= 10
		}
		if tmp == l {
			c++

		}
		counter++
		//check first maxLenght digit is same with l
	}
	fmt.Println("res", c)
	fmt.Println("res", counter)
	return 0
}
