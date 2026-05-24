package solution

import "slices"

func threeSum(nums []int) [][]int {
	slices.Sort(nums)
	n := len(nums)
	res := make([][]int, 0)
	for i, num := range nums[:n-2] {
		if i > 0 && num == nums[i - 1] {
			continue
		}
		if num + nums[i+1] + nums[i+2] > 0 {
			break
		}
		if num + nums[n-1] + nums[n-2] < 0 {
			continue
		}
		j, k := i + 1, n - 1
		for j < k {
			curSum := num + nums[j] + nums[k]
			if curSum < 0 {
				j++
			} else if curSum > 0 {
				k--
			} else {
				if j == i + 1 || nums[j] != nums[j - 1] { // 重要
					res = append(res, []int{num, nums[j], nums[k]})
				}
				j++
				k--
			}
		}
	}
	return res
}
