package problem

import "fmt"

func Problem6() {
	// <p>The sum of the squares of the first ten natural numbers is,</p>
	// $$1^2 + 2^2 + ... + 10^2 = 385$$
	// <p>The square of the sum of the first ten natural numbers is,</p>
	// $$(1 + 2 + ... + 10)^2 = 55^2 = 3025$$
	// <p>Hence the difference between the sum of the squares of the first ten natural numbers and the square of the sum is $3025 - 385 = 2640$.</p>
	// <p>Find the difference between the sum of the squares of the first one hundred natural numbers and the square of the sum.</p>
	//1+2+3...+10 = 55
	//(1+10)*(10/2)

	problem6fn()
}

func problem6fn() {
	r1 := 0
	r2 := 0
	for i := 1; i <= 100; i++ {
		r1 += i * i
		r2 += i
	}
	fmt.Println((r2 * r2) - r1)
}
