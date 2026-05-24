package solution

func getMinDistance(nums []int, target int, start int) int {
	n := len(nums)
	i, j := start, start
	for i >= 0 || j < n {
		if i >= 0 {
			if nums[i] == target {
				return start - i
			}
			i--
		}
		if j < n {
			if nums[j] == target {
				return j - start
			}
			j++
		}
	}
	return -1
}
