package problem

import (
	"fmt"
	"math"
	"math/big"
)

func Problem66() {
	problem66fn()
}

func problem66fn() {
	//https://en.wikipedia.org/wiki/Pell%27s_equation
	maxX := big.NewInt(0)
	maxI := 0
	for i := 2; i < 1000; i++ {
		if math.Sqrt(float64(i)) != float64(int(math.Sqrt(float64(i)))) {
			fmt.Println("Finding i", i)
			x, _ := pellEq(i)
			if maxX.Cmp(x) < 0 {
				maxX = x
				maxI = i
			}
		}
	}
	fmt.Println(maxX, maxI)

}

// return x,y
func pellEq(i int) (*big.Int, *big.Int) {
	bi := big.NewInt(int64(i))
	ll := int64(math.Sqrt(float64(i)))
	lim := big.NewInt(ll)
	m := big.NewInt(0)
	d := big.NewInt(1)
	a := big.NewInt(ll)

	h := big.NewInt(int64(ll))
	h1 := big.NewInt(1)
	k := big.NewInt(1)
	k1 := big.NewInt(0)
	voila := false
	neg := big.NewInt(int64(-1))
	one := big.NewInt(int64(1))
	for !voila {
		pellApprox := new(big.Int).Add(new(big.Int).Mul(h, h), new(big.Int).Mul(new(big.Int).Mul(neg, bi), new(big.Int).Mul(k, k)))
		fmt.Println("calc", i, h, k, pellApprox)
		if pellApprox.Cmp(one) == 0 {
			voila = true
		} else {
			//Update the h & k based on https://en.wikipedia.org/wiki/Continued_fraction
			//https://en.wikipedia.org/wiki/Continued_fraction
			m = m.Add(new(big.Int).Mul(d, a), new(big.Int).Mul(neg, m))
			d = d.Div(new(big.Int).Add(bi, new(big.Int).Mul(neg, new(big.Int).Mul(m, m))), d)
			a = a.Div(new(big.Int).Add(lim, m), d)
			h2 := h1
			h1 = h
			k2 := k1
			k1 = k
			h = new(big.Int).Add(h2, new(big.Int).Mul(a, h1))
			k = new(big.Int).Add(k2, new(big.Int).Mul(a, k1))
		}
	}
	return h, k
}
