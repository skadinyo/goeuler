package problem

import "fmt"

func Problem2() {
	result := problem2fn(1, 1, 4000000)
	fmt.Println(result)
}

func problem2fn(a, b, max int) int {
	i := a
	j := b
	k := 0
	res := 0
	for k < max {
		k = i + j
		if k%2 == 0 {
			res += +k
		}
		i = j
		j = k
	}
	return res
}
