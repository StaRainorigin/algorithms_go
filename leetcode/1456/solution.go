package solution

func maxVowels(s string, k int) int {
	isVowel := func(c byte) bool {
		return c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u'
	}
	n := len(s)
	count := 0
	for i := range k {
		if isVowel(s[i]) {
			count++
		}
	}
	ans := count
	for i := k; i < n; i++ {
		if isVowel(s[i]) {
			count++
		}
		if isVowel(s[i-k]) {
			count--
		}
		ans = max(ans, count)
	}
	return ans
}
