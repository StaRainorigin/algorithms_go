package solution

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	
	cntS, cntT := [256]byte{}, [256]byte{}
	
	for i := range s {
		c1, c2 := s[i], t[i]
		
		if cntS[c1] != 0 && cntS[c1] != c2 {
			return false
		}
		if cntT[c2] != 0 && cntT[c2] != c1 {
			return false
		}
		
		cntS[c1] = c2
		cntT[c2] = c1
	}
	
	return true
}
