package main

import (
	problem "github.com/skadinyo/goeuler/problem"
)

func main() {
	// TODO parse it from argument
	Solve(3)
}

func Solve(i int) {
	// TODO find a way to make it meta
	fnMap := map[int]func(){
		1: problem.Problem1,
		2: problem.Problem2,
		3: problem.Problem3,
	}
	fnMap[i]()
}
