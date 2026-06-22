package solution

func maxNumberOfBalloons(text string) int {
	cnt := [26]int{}
	for _, c := range text {
		cnt[c-'a']++
	}
	return min(cnt[0], cnt[1], cnt[11]/2, cnt[13], cnt[14]/2)
}

// func maxNumberOfBalloons(text string) int {
// 	cnt := [5]int{}
// 	mp := map[rune]int{
// 		'b': 0,
// 		'a': 1,
// 		'l': 2,
// 		'o': 3,
// 		'n': 4,
// 	}

// 	for _, c := range text {
// 		if i, ok := mp[c]; ok {
// 			cnt[i]++
// 		}
// 	}

// 	cnt[2] /= 2
// 	cnt[3] /= 2

// 	ans := 10_000
// 	for _,x := range cnt {
// 		ans = min(ans, x)
// 	}

// 	if ans == 10_000 {
// 		return 0
// 	}
// 	return ans
// }
