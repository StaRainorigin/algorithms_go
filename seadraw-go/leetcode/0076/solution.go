package solution

func minWindow(s string, t string) string {
	n, m := len(s), len(t)
	if m == 0 || m > n {
		return ""
	}

	cntS := [128]bool{}
	cntT := [128]int{}

	less := 0
	for _, c := range t {
		if cntT[c] == 0 {
			less++
			cntS[c] = true
		}
		cntT[c]++
	}

	l, r := 0, n+1
	i := 0
	for j, c := range s{
		if !cntS[c] {
			continue
		}

		cntT[c]--
		if cntT[c] == 0 {
			less--
		}

		for i <= j && (!cntS[s[i]] || cntT[s[i]] < 0) {
			if cntS[s[i]] {
				cntT[s[i]]++
			}
			i++
		}

		if less == 0 && j+1-i < r-l {
			l, r = i, j+1
		}
	}

	if r == n+1 {
		return ""
	}
	return s[l:r]
}
