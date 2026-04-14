package solution

func productExceptSelf(nums []int) []int {
	n := len(nums)
	leftMults, rightMults := make([]int, n), make([]int, n)
	leftMults[0], rightMults[n - 1] = 1, 1
	for i := range n - 1 {
		leftMults[i + 1] = leftMults[i] * nums[i]
		rightMults[n - i - 2] = rightMults[n - i - 1] * nums[n - i - 1]
	}
	for i := range n {
		leftMults[i] *= rightMults[i]
	}
	return leftMults
}
