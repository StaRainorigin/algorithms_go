package solution

func minMirrorPairDistance(nums []int) int {
	n := len(nums)
	ans := n
	mp := make(map[int]int)
	for idx, num := range nums {
		if pre, ok := mp[num]; ok {
			ans = min(ans, idx - pre)
		}
		
		rev := 0
		for ; num > 0; num /= 10 {
			rev = rev * 10 + num % 10
		}
		mp[rev] = idx
	}
	
	if ans == n {
		ans = -1
	}
	return ans
}
