package problem

import (
	"fmt"
)

func Problem1() {
	problem1fn(3, 5, 1000)
}

func problem1fn(a int, b int, max int) int {
	result := 0
	for i := 1; i < max; i++ {
		if i%a == 0 || i%b == 0 {
			result += +i
		}
	}
	fmt.Println(result)
	return result
}
