package solution

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func solveQueries(nums []int, queries []int) []int {
	n := len(nums)

	pos := make(map[int][]int, n)
	for i, x := range nums {
		pos[x] = append(pos[x], i)
	}

	best := make([]int, n)
	for i := range best {
		best[i] = -1
	}

	for _, idxs := range pos {
		m := len(idxs)
		if m == 1 {
			continue
		}

		for i, idx := range idxs {
			prev := idxs[(i-1+m)%m]
			next := idxs[(i+1)%m]

			left := idx - prev
			if left < 0 {
				left += n
			}

			right := next - idx
			if right < 0 {
				right += n
			}

			best[idx] = min(left, right)
		}
	}

	ans := make([]int, len(queries))
	for i, q := range queries {
		ans[i] = best[q]
	}
	return ans
}


// import (
// 	"sort"
// )

// func abs(x int) int {
// 	if x < 0 {
// 		return -x
// 	}
// 	return x
// }

// func solveQueries(nums []int, queries []int) []int {
// 	n := len(nums)
// 	pos := make(map[int][]int)
// 	for i, x := range nums {
// 		pos[x] = append(pos[x], i)
// 	}
	
// 	ans := make([]int, len(queries))
// 	for i, q := range queries {
// 		indices := pos[nums[q]]
// 		if len(indices) == 1 {
// 			ans[i] = -1
// 			continue
// 		}
		
// 		idx := sort.SearchInts(indices, q)
// 		m := len(indices)
// 		pre := indices[(idx-1+m)%m]
// 		nxt := indices[(idx+1)%m]
		
// 		d1 := abs(q - pre)
// 		d2 := abs(q - nxt)
// 		ans[i] = min(min(d1, n-d1), min(d2, n-d2))
// 	}
// 	return ans
// }

// 超时
// func solveQueries(nums []int, queries []int) []int {
// 	n := len(nums)
// 	ans := make([]int, len(queries))
// 	for i, startIndex := range queries {
// 		cur := nums[startIndex]
// 		res := -1
// 		for j := 1; j < n/2+1; j++ {
// 			if nums[(startIndex+j)%n] == cur || nums[(startIndex+n-j)%n] == cur {
// 				res = j
// 				break
// 			}
// 		}
// 		ans[i] = res
// 	}
// 	return ans
// }
