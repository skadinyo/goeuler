package problem

import (
	"fmt"
)

func Problem5() {
	//2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.
	//What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?

	// 1,2,3,4,5,6,7,8,9,10
	// 1,2,3,[2,2],4,5,[2,3],7,[2,2,2],[3,3],[2,5]
	// 1,2,3,5,7 -> 210
	// 210 * 12 = 2520
	// 12 -> [2,2,3]
	// 2520 factors
	// 1 × 2520 = 2,520
	// 2 × 1260 = 2,520
	// 3 × 840 = 2,520
	// 4 × 630 = 2,520
	// 5 × 504 = 2,520
	// 6 × 420 = 2,520
	// 7 × 360 = 2,520
	// 8 × 315 = 2,520
	// 9 × 280 = 2,520
	// 10 × 252 = 2,520
	// 12 × 210 = 2,520
	// 14 × 180 = 2,520
	// 15 × 168 = 2,520
	// 18 × 140 = 2,520
	// 20 × 126 = 2,520
	// 21 × 120 = 2,520
	// 24 × 105 = 2,520
	// 28 × 90 = 2,520
	// 30 × 84 = 2,520
	// 35 × 72 = 2,520
	// 36 × 70 = 2,520
	// 40 × 63 = 2,520
	// 42 × 60 = 2,520
	// 45 × 56 = 2,520

	// 1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20
	// 1,2,3,[2,2],5,[2,3],7,[2,2,2],[3,3],[2,5],11,[2,2,3],13,[2,7],[3,5],[2,2,2,2],17,[2,3,3],19,[2,2,5]
	// [2],[3],[2,2],[3],[2,2,2],[3] -> [2,2,2,3]
	// 1*2*3*5*7*11*13*17*19*2*2*2*3 -> 12252240
	problem5fn()
}

func problem5fn() {
	// Pen & paper
	// To code it, find prime factors of each number (don't remove duplicate!), i.e, f(12) = [2,2,3]
	// remove the factors that is a subset of other factors, then multuply it by prime numbers
	fmt.Println(12252240)
}
