package solution

func checkInclusion(s1 string, s2 string) bool {
	m, n := len(s1), len(s2)
	if m > n {
		return false
	}
	
	aMap := [26]int{}
	for _, c := range s1 {
		aMap[c-'a']++
	}

	bMap := [26]int{}
	for i, c := range s2 {
		bMap[c-'a']++
		if i < m - 1 {
			continue
		}
		
		if aMap == bMap {
			return true
		}
		
		bMap[s2[i-m+1]-'a']--
	}
	
	return false
}
