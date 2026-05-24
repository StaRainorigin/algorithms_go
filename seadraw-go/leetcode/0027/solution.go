package solution

func removeElement(nums []int, val int) int {
	n := len(nums)
	res := n
	i, j := 0, 0
	for j < n {
		if nums[j] != val {
			if i != j {
				nums[i] = nums[j]
			}
			i++
		} else {
			res--
		}
		j++
	}
	return res
}

