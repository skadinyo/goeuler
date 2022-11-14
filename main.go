package main

import (
	problem "github.com/skadinyo/goeuler/problem"
)

func main() {
	// TODO parse it from argument
	Solve(66)
}

func Solve(i int) {
	// TODO find a way to make it meta
	fnMap := map[int]func(){
		1:   problem.Problem1,
		2:   problem.Problem2,
		3:   problem.Problem3,
		4:   problem.Problem4,
		5:   problem.Problem5,
		6:   problem.Problem6,
		7:   problem.Problem7,
		8:   problem.Problem8,
		9:   problem.Problem9,
		10:  problem.Problem10,
		719: problem.Problem719,
		684: problem.Problem684,
		66:  problem.Problem66,
	}

	fnMap[i]()
}

func GenTemplate() {
	// TODO
	// create files for problem x
}
