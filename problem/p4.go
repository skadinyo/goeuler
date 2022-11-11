package problem

import (
	"fmt"

	lib "github.com/skadinyo/goeuler/lib"
)

func Problem4() {
	//A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.
	//Find the largest palindrome made from the product of two 3-digit numbers.
	problem4fn()
}

// palindrome
// 11
// 111
// 121
// 131
// ...
// 212
// 222
// ...
// 999
// 1001
// ....
// 10001
// .....
/// 998899
/// 999999
// 1000001 -> max

//

// three digit number means max number is 1000*1000 or 1.000.000
// and min is 100*100 or 10.000
// brute force need to do ~(1000-100)C2 -> ~404550  calc.
func problem4fn() int {
	// Brute
	// dont even care about the same repetitions
	result := 0
	for i := 100; i < 1000; i++ {
		for j := 100; j < 1000; j++ {
			k := i * j
			if lib.IsPalin(k) && (result < k) {
				result = k
			}
		}
	}
	fmt.Println(result)
	return result
}
