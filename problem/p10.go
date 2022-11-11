package problem

import "fmt"

func Problem10() {
	problem10fn()
}

// Sieveing prime
func problem10fn() {
	max := 2000000
	xs := []bool{}
	for i := 0; i <= max; i++ {
		xs = append(xs, true)
	}
	xs[0] = false
	xs[1] = false
	i := 2
	// Should hav euse sqrt for the limit, but meh
	for i <= ((max / 2) + 1) {
		if xs[i] {
			// Remove multiples
			for j := i + i; j <= max; j += i {
				xs[j] = false
			}
		}
		i++
	}
	res := 0
	for k, v := range xs {
		if v {
			res += k
		}
	}
	fmt.Println(res)
}
