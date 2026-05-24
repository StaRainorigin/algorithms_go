package solution

func search(nums []int, target int) int {
	l, r := 0, len(nums)-1

	for l <= r {
		m := (l + r) / 2
		if nums[m] == target {
			return m
		}

		// 左半边有序
		if nums[l] <= nums[m] {
			if nums[l] <= target && target < nums[m] {
				r = m - 1  // target 在左边
			} else {
				l = m + 1  // target 在右边
			}
		// 右半边有序
		} else {
			if nums[m] < target && target <= nums[r] {
				l = m + 1  // target 在右边
			} else {
				r = m - 1  // target 在左边
			}
		}
	}

	return -1
}
