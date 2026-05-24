package solution

func twoSum(numbers []int, target int) []int {
	i, j := 0, len(numbers) - 1
	res := []int{-1, -1}
	for i < j {
		cur := numbers[i] + numbers[j]
		if cur > target {
			j--
		} else if cur < target {
			i++
		} else {
			res[0] = i + 1
			res[1] = j + 1
			break
		}
	}
	return res
}
