package solution

func canReach(s string, minJump, maxJump int) bool {
	n := len(s)
	canReaches := make([]bool, n)
	canReaches[0] = true
	j := 1
	for i, ch := range s {
		if ch == '0' && canReaches[i] {
			// 注意 j 只会增大，不会减小，所以总体时间复杂度是 O(n)
			for j = max(j, i+minJump); j <= min(i+maxJump, n-1); j++ {
				canReaches[j] = true // 可以跳到 j
			}
		}
	}
	return s[n-1] == '0' && canReaches[n-1]
}