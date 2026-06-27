package solution

func countMajoritySubarrays(nums []int, target int) int {
	i := 0
	ans := 0
	count := 0
	for j := range nums {
		if nums[j] == target {
			count++
		}
		if count > (j - i) / 2 {
			ans++
			continue
		}
		for i <= j && count <= (j - i) / 2 {
			if nums[i] == target {
				count--
			}
			i++
		}
	}
	return ans
}

