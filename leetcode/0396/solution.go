package solution

func maxRotateFunction(nums []int) int {
	n := len(nums)
	sum, f := 0, 0
	for i, num := range nums {
		sum += num
		f += num * i
	}
	ans := f
	for i := range n {
		f += sum - n * nums[n-i-1]
		ans = max(ans, f)
	}
	return ans
}
