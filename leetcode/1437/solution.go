package solution

func kLengthApart(nums []int, k int) bool {
	count := k
	res := true
	for _, num := range nums {
		if num == 1 {
			if count < k {
				res = false
				break
			}
			count = 0
		} else {
			count++
		}
	}
	return res
}
