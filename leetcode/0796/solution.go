package solution

func rotateString(s string, goal string) bool {
	n := len(s)
	ss := s + s
	res := false
	for i := range n {
		if ss[i:i+n] == goal {
			res = true
			break
		}
	}
	return res
}
