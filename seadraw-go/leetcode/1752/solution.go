package solution

func check(nums []int) bool {
	sorted := true
	for i := 1; i < len(nums); i++ {
		if nums[i-1] > nums[i] {
			if !sorted {
				return false
			}
			sorted = false
		}
	}
	return sorted || nums[0] >= nums[len(nums)-1]
}
