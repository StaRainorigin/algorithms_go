package solution

func numberOfSpecialChars(word string) int {
	cnt := [26]int{}
	for _, c := range word {
		if 'a' <= c && c <= 'z' && cnt[c - 'a'] == 0 {
			cnt[c - 'a'] = 1
		}
		if 'A' <= c && c <= 'Z' && cnt[c - 'A'] == 1 {
			cnt[c - 'A'] = 2
		}
		if 'A' <= c && c <= 'Z' && cnt[c - 'A'] == 0 {
			cnt[c - 'A'] = 3
		}
		if 'a' <= c && c <= 'z' && cnt[c - 'a'] == 2 {
			cnt[c - 'a'] = 3
		}
 	}

	ans := 0
	for _, x := range cnt {
		if x == 2 {
			ans++
		}
	}

	return ans
}
