package solution

// 4 4 4 2 3 4 4 4 4 11 => 5

func maximumElementAfterDecrementingAndRearranging(arr []int) int {
	n := len(arr)
	cnt := make([]int, n + 1)
	for _, x := range arr {
		cnt[min(x, n)]++
	}

	ans := 0
	for i := 1; i <= n; i++ {
		ans = min(ans + cnt[i], i)
	}

	return ans
}
