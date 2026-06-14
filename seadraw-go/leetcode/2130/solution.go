package solution

type ListNode struct {
	Val  int
	Next *ListNode
}

func pairSum(head *ListNode) int {
	nums := []int{}
	
	for p := head; p != nil; p = p.Next {
		nums = append(nums, p.Val)
	}

	ans := 0
	n := len(nums)
	for i := range n/2 {
		ans = max(ans, nums[i] + nums[n-1-i])
	}

	return ans
}
