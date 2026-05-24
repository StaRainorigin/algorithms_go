package solution

import "strconv"

func summaryRanges(nums []int) []string {
	n := len(nums)
	ans := make([]string, 0)
	i := 0

	for j := range n {
		if j == n-1 || nums[j+1] != nums[j]+1 {
			if i == j {
				ans = append(ans, strconv.Itoa(nums[i]))
			} else {
				ans = append(ans, strconv.Itoa(nums[i])+"->"+strconv.Itoa(nums[j]))
			}
			i = j + 1
		}
	}
	return ans
}