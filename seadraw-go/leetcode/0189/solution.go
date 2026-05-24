package solution

func reverse(nums []int, left int, right int) {
	for left < right {
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}
}

func rotate(nums []int, k int)  {
	n := len(nums)
	
	if n <= 1 {
		return
	}
	
	k %= n
	
	reverse(nums, 0, n - 1)
	reverse(nums, 0, k - 1)
	reverse(nums, k, n - 1)
	// left1, right1 := 0, n - k - 1
	// left2, right2 := n - k, n - 1
	// reverse(nums, left1, right1)
	// reverse(nums, left2, right2)
	// reverse(nums, 0, n - 1)
}