package problem

import (
	"fmt"

	lib "github.com/skadinyo/goeuler/lib"
)

func Problem719() {
	problem719fn()
}

func problem719fn() {

	r := 0
	for i := 1; i <= 1000000; i++ {
		sq := i * i
		sqDigs := lib.IntToArray(sq)
		poss := splitNumber(i, sqDigs, (len(sqDigs)/2)+1)
		if checkNatural(i, poss) {
			r += sq
			fmt.Println(i, sq)
		}
	}
	fmt.Println(r)

}

func splitNumber(s int, arr []int, maxLength int) [][][]int {
	l := len(arr)
	if l <= 1 {
		return [][][]int{}
	}
	res := [][][]int{[][]int{[]int{arr[0]}}} // initial

	for i := 1; i < l; i++ {
		tmpRes := [][][]int{}
		removeSlice := []int{}
		// Playing with slice pointer is painful :\
		for x, xs := range res {
			// append new
			tmpXs := [][]int{}
			for _, ys := range xs {
				tmpXs = append(tmpXs, ys)
			}
			tmpXs = append(tmpXs, []int{arr[i]})
			tmpRes = append(tmpRes, tmpXs)
			// modify current xs with arr[i]
			ys := xs[len(xs)-1]
			if len(ys) <= maxLength || lib.ArrayToInt(ys) < i {
				ys = append(ys, arr[i])
				xs[len(xs)-1] = ys
				res[x] = xs
			} else {
				removeSlice = append(removeSlice, x)
			}

		}
		res = append(res, tmpRes...)
		for _, x := range removeSlice {
			res = append(res[:x], res[x+1:]...)
		}
	}
	return res
}

func splitNumberFn(sourceArr []int, i int, l int, target [][][]int) [][][]int {

	if i == l {
		return target
	}

	top := splitNumberFn(sourceArr, i+1, l, [][][]int{})
	bot := splitNumberFn(sourceArr, i+1, l, [][][]int{})

	target = append(target, top...)
	target = append(target, bot...)
	return target
}

func checkNatural(x int, arr [][][]int) bool {
	for _, arr2 := range arr {
		y := 0
		for _, arr3 := range arr2 {
			y += lib.ArrayToInt(arr3)
		}
		if x == y {
			fmt.Println(arr2)
			return true
		}
	}
	return false
}

// arr: Store input array
// i: Stores current index of arr
// N: Stores length of arr
// K: Stores count of subsets
// nos: Stores count of feasible subsets formed
// v: Store K subsets of the given array
// 123
// k 2
// [1 2 3], 0, 3, 2, 0, [[] []]
//
// j Looping 0,1
// masuk else semua karena v[j] kosong semua
// [1 2 3], 1, 3, 2, 1, [[1] []]
// j looping 0,1
// 0 -> masuk len(v[j]) > 0
// [1 2 3], 2, 3, 2, 1, [[1 2] []] -> [1 2 3], 3, 3, 2, 1, [[1 2 3] []] && [1 2 3], 3, 3, 2, 2, [[1 2] [3]]
// v balik ke [[1] []]
// 1 -> masuk else
// [1 2 3], 2, 3, 2, 2, [[1] [2]] -> [1 2 3], 3, 3, 2, 2, [[1 3] [2]] && [1 2 3], 3, 3, 2, 2, [[1] [2 3]]
//
func partitionSub(arr []int, i, n, k, nos int, v [][]int) [][][]int {
	result := [][][]int{}
	if i >= n {
		if nos == k {
			r := [][]int{}
			for x := 0; x < len(v); x++ {
				s := []int{}
				for y := 0; y < len(v[x]); y++ {
					s = append(s, v[x][y])
				}
				r = append(r, s)
			}
			result = append(result, r)
			return result
		}
		return nil
	}
	for j := 0; j < k; j++ {
		if len(v[j]) > 0 {
			v[j] = append(v[j], arr[i])
			r1 := partitionSub(arr, i+1, n, k, nos, v)
			if r1 != nil {
				result = append(result, r1...)
			}
			v[j] = v[j][:len(v[j])-1]
		} else {
			v[j] = append(v[j], arr[i])
			r2 := partitionSub(arr, i+1, n, k, nos+1, v)
			if r2 != nil {
				result = append(result, r2...)
			}
			v[j] = v[j][:len(v[j])-1]
			break
		}
	}
	return result
}

func partitionKSub(arr []int, n, k int) [][][]int {

	v := [][]int{}
	i := 0
	for i < k {
		v = append(v, []int{})
		i++
	}

	if (k == 0) || (k > n) {
		return [][][]int{}
	}
	return partitionSub(arr, 0, n, k, 0, v)

}
