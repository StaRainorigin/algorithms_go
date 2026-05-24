package solution

func twoSum(nums []int, target int) []int {
	cnt := make(map[int]int)
	for i, num := range nums {
		if j, ok := cnt[target - num]; ok {
			return []int{i, j}
		}
		cnt[num] = i
	}
	return []int{-1, -1}
}
