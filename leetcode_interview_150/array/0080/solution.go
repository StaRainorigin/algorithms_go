package solution

func removeDuplicates(nums []int) int {
	n := len(nums)
	if n <= 2 {
		return n
	}

	write := 2
	for read := 2; read < n; read++ {
		if nums[read] != nums[write-2] {
			nums[write] = nums[read]
			write++
		}
	}
	return write
}

// 1, 2, 2, 3, 4, 4, 4, 5, 5, 5, 5, 5, 6, 6