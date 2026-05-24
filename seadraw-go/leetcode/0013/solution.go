package soultion

var maps = map[rune]int {
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func romanToInt(s string) int {
	res := 0
	n := len(s)
	for i := 0; i < n - 1; i++ {
		cur := maps[rune(s[i])]
		nxt := maps[rune(s[i+1])]
		if cur < nxt {
			res -= cur
		} else {
			res += cur
		}
	}
	res += maps[rune(s[n-1])]
	return res
}
