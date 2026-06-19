package solution

func largestAltitude(gain []int) int {
	ans := 0
	cur := 0
	for _, x := range gain {
		cur += x
		ans = max(ans, cur)
	}
	return ans
}
