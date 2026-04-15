package solution

func hIndex(citations []int) int {
	n := len(citations)
	res := 0
	counts := make([]int, n + 1)
	for _, c := range citations {
		if c < n + 1 {
			counts[c]++
		} else {
			counts[n]++
		}
	}
	sum := 0
	for i := n; i > 0; i-- {
		sum += counts[i]
		if sum >= i {
			res = i
			break
		}
	}
	return res
}