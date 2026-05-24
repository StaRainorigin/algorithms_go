package solution

func isValid(s string) bool {
	st := []byte{}
	m := map[rune]byte {
		'(': ')',
		'[': ']',
		'{': '}',
	}
	for _, c := range s {
		n := len(st)
		d, ok:= m[c] 
		if ok {
			st = append(st, d)
			n++
		} else {
			if n <= 0 {
				return false
			}
			if c == rune(st[n-1]) {
				st = st[:n-1]
				n--
			} else {
				return false
			}
		}
	}
	if len(st) != 0 {
		return false
	}
	return true
}
