package problem

import "fmt"

func Problem3() {
	//The prime factors of 13195 are 5, 7, 13 and 29.
	//What is the largest prime factor of the number 600851475143 ?
	result := problem3fn(600851475143)
	fmt.Println(result)
}

func problem3fn(x int) int {
	i := 2
	max := x
	for max != 1 {
		for max%i != 0 {
			i++
		}
		max /= i
	}
	return i
}
