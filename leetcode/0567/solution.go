package solution

func checkInclusion(s1 string, s2 string) bool {
	m, n := len(s1), len(s2)
	if m > n {
		return false
	}
	
	less := 0
	cnt := [26]int{}
	for _, c := range s1 {
		if cnt[c - 'a'] == 0 {
			less++
		}
		cnt[c - 'a']++
	}
	
	for r, c := range s2 {
		cnt[c - 'a']--
		if cnt[c - 'a'] == 0 {
			less--
		}
		
		if r < m - 1 {
			continue
		}
		
		if less == 0 {
			return true
		}
		
		l := r-m+1
		if cnt[s2[l] - 'a'] == 0 {
			less++
		}
		cnt[s2[l] - 'a']++
	}
	return false
}
