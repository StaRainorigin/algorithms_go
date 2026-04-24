package solution

func abs(x int64) int64 {
	if x < 0 {
		return -x
	}
	return x
}

func distance(nums []int) []int64 {
	ans := make([]int64, len(nums))
	cnt := map[int][]int{}
	for i, num := range nums {
		cnt[num] = append(cnt[num], i)
	}
	for _, arr := range cnt {
		n := len(arr)
		if n == 0 {
			continue
		}
		sum := int64(0)
		for _, idx := range arr {
			sum += abs(int64(arr[0]) - int64(idx))
		}
		ans[arr[0]] = sum
		for i := 1; i < n; i++ {
			gap := int64(arr[i] - arr[i-1])
			sum += int64(i) * gap
			sum -= int64(n - i) * gap
			ans[arr[i]] = sum
		}
	}
	return ans
}

// func distance(nums []int) []int64 {
// 	ans := make([]int64, len(nums))
// 	cnt := map[int][]int{}
// 	for i, num := range nums {
// 		cnt[num] = append(cnt[num], i)
// 	}
// 	for _, arr := range cnt {
// 		n := len(arr)
// 		sum := make([]int, n + 1)
// 		for i, x := range arr {
// 			sum[i+1] = sum[i] + x
// 		}
// 		for i, x := range arr {
// 			left := x * i - sum[i]
// 			right := sum[n] - sum[i] - x * (n - i)
// 			ans[x] = int64(left + right)
// 		}
// 	}
// 	return ans
// }
