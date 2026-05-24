package solution

func isSubsequence(s string, t string) bool {
	m, n := len(s), len(t)
	if m > n {
		return false
	}
	if m == 0 {
		return true
	}
	i := 0
	for _, c := range t {
		if rune(s[i]) == c {
			i++
			if i == m {
				return true
			}
		}
	}
	return false
}
