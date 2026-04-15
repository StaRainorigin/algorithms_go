package solution

func lengthOfLongestSubstring(s string) int {
	i, j, n := 0, 0, len(s)
	res := 0
	aMap := make(map[byte]int)
	for j < n {
		idx, ok := aMap[s[j]]
		if ok {
			res = max(res, j-i)
			for i <= idx {
				delete(aMap, s[i])
				i++
			}
		} else {
			aMap[s[j]] = j
			j++
		}
	}
	res = max(res, j-i)
	return res
}
