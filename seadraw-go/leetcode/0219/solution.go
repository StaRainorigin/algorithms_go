package solution

func containsNearbyDuplicate(nums []int, k int) bool {
	cnt := map[int]int{}
	for i, num := range nums {
		if idx, ok := cnt[num]; ok && i - idx <= k {
			return true
		}
		cnt[num] = i
	}
	return false
}
