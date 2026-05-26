package solution

import "math/bits"

func numberOfSpecialChars(word string) int {
	mask := [2]int{}
	for _, c := range word {
		mask[c>>5&1] |= 1 << (c & 31)
	}
	return bits.OnesCount(uint(mask[0] & mask[1]))
}

// func numberOfSpecialChars(word string) int {
// 	cnt := [52]bool{}
// 	ans := 0
// 	for _, c := range word {
// 		if 'a' <= c && c <= 'z' {
// 			cnt[c - 'a'] = true
// 		}
// 		if 'A' <= c && c <= 'Z' {
// 			cnt[c - 'A' + 26] = true
// 		}
// 	}

// 	for i := range 26 {
// 		if cnt[i] && cnt[i+26] {
// 			ans++
// 		}
// 	}

// 	return ans
// }
