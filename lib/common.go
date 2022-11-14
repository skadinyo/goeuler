package lib

func LeastCommonMultiple(a, b int) int {
	return a * b / GreatestCommonDivisor(a, b)

}

//https://rosettacode.org/wiki/Greatest_common_divisor#Go
func GreatestCommonDivisor(a, b int) int {
	var bgcd func(a, b, res int) int

	bgcd = func(a, b, res int) int {
		switch {
		case a == b:
			return res * a
		case a%2 == 0 && b%2 == 0:
			return bgcd(a/2, b/2, 2*res)
		case a%2 == 0:
			return bgcd(a/2, b, res)
		case b%2 == 0:
			return bgcd(a, b/2, res)
		case a > b:
			return bgcd(a-b, b, res)
		default:
			return bgcd(a, b-a, res)
		}
	}

	return bgcd(a, b, 1)
}

func PrimeSieve(max int) []bool {
	return nil
}

func IsPrimeDumb(x int) bool {
	if x == 1 {
		return false
	}
	for i := 2; i <= ((x / 2) + 1); i++ {
		if x%i == 0 {
			return false
		}
	}
	return true
}

func IntToArray(i int) []int {
	res := []int{}
	x := i
	if x == 0 {
		return []int{0}
	}
	for x >= 10 {
		rem := x % 10
		res = append([]int{rem}, res...)
		x /= 10
	}
	res = append([]int{x}, res...)
	return res
}

func ArrayToInt(arr []int) int {
	res := 0
	max := len(arr) - 1
	i := 0
	m := 1
	for i <= max {
		res += arr[max-i] * m
		m *= 10
		i++
	}
	return res
}

func Max(xs []int) int {
	res := 0
	for i := range xs {
		if res < i {
			res = i
		}
	}
	return res
}

func IsPalin(x int) bool {
	digit := IntToArray(x)
	l := len(digit)
	for i := 0; i <= l/2; i++ {
		if digit[i] != digit[l-i-1] {
			return false
		}
	}
	return true
}
