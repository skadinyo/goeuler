package problem

import (
	"fmt"
)

func Problem5() {
	problem5fn()
}

func problem5fn() int {
	result := 1
	for i := 1; i <= 10; i++ {
		result *= i
	}
	fmt.Println(result)
	return result
}
