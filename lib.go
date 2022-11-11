package main

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

func PrimeSieve(max int) []int {
	return nil
}

func IntToArray(i int) []int {
	res := []int{}
	x := 1
	if x == 0 {
		return []int{0}
	}
	for x != 0 {
		rem := x % 10
		res = append(res, rem)
		x /= 10
	}
	return res
}
