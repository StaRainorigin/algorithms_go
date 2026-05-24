package solution

func lengthOfLongestSubstring(s string) (ans int) {
	window := [128]bool{}
	i := 0
	for j, c := range s {
		for window[c] {
			window[s[i]] = false
			i++
		}
		window[c] = true
		ans = max(ans, j - i + 1)
	}
	return
}
