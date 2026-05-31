package solution

import (
	"math"
	"math/bits"
	"slices"
)

func asteroidsDestroyed(mass int, asteroids []int) bool {
	mWid := bits.Len(uint(slices.Max(asteroids)))
	sums := make([]int, mWid)
	mins := make([]int, mWid)
	for i := range mins {
		mins[i] = math.MaxInt
	}

	for _, x := range asteroids {
		i := bits.Len(uint(x)) - 1
		sums[i] += x
		mins[i] = min(mins[i], x)
	}

	res := true
	for i, x := range mins {
		if x == math.MaxInt {
			continue
		}

		if mass < x {
			res = false
			break
		} 

		mass += sums[i]
	}

	return res
}


// func asteroidsDestroyed(mass int, asteroids []int) bool {
// 	slices.Sort(asteroids)
// 	cur := mass
// 	res := true
// 	for _, x := range asteroids {
// 		if cur < x {
// 			res = false
// 			break
// 		} else {
// 			cur += x
// 		}
// 	}
// 	return res
// }
