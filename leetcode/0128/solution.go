package solution

func longestConsecutive(nums []int) int {
	has := map[int]bool{}
	for _, num := range nums {
		has[num] = true
	}
	
	ans := 0
	for x, _ := range has {
		if has[x - 1] {
			continue
		}
		
		y := x + 1
		for has[y] {
			y++
		}
		ans = max(ans, y - x)
	}
	return ans
}
