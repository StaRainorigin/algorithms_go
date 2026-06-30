package solution

func numberOfSubstrings(s string) int {
	ans := 0
	cnt := [3]int{}
	i := 0
	for _, c := range s {
		cnt[c-'a']++
		for cnt[0] > 0 && cnt[1] > 0 && cnt[2] > 0 {
			cnt[s[i]-'a']--
			i++
		}
		ans += i
	}
	return ans
}
