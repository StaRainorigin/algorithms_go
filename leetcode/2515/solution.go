package solution

func closestTarget(words []string, target string, startIndex int) int {
	n := len(words)
	res := -1
	for i := range n/2 + 1 {
		if words[(startIndex + i) % n] == target || words[(startIndex + n - i) % n] == target{
			res = i 
			break
		}
	}
	return res
}
