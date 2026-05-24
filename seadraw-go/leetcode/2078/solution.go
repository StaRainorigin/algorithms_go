package solution

func maxDistance(colors []int) int {
	n := len(colors)
	ans := -1
	if n == 0 {
		return ans
	}
	l, r := colors[0], colors[n-1]
	if l != r {
		return n - 1
	}
	for i := 1; i < n - 1; i++ {
		if colors[i] != l {
			ans = max(ans, i - 0)
		}
		if colors[i] != r {
			ans = max(ans, n - 1 - i)
		}
	}
	return ans
}
