package solution

import "math/bits"

func findThePrefixCommonArray(A []int, B []int) []int {
	var p, q uint
	for i := range A {
		p |= 1 << A[i]
		q |= 1 << B[i]
		A[i] = bits.OnesCount(p & q)
	}
	return A
}

// func findThePrefixCommonArray(A []int, B []int) []int {
// 	mp := map[int]bool{}
// 	n := len(A)
// 	count := 0
// 	ans := make([]int, n)
// 	for i := range n {
// 		x := A[i]
// 		_, ok := mp[x]
// 		if ok {
// 			count++
// 		} else {
// 			mp[x] = true
// 		}

// 		x = B[i]
// 		_, ok = mp[x]
// 		if ok {
// 			count++
// 		} else {
// 			mp[x] = true
// 		}

// 		ans[i] = count
// 	}
// 	return ans
// }
