package solution

import "strings"

func wordPattern(pattern string, s string) bool {
	t := strings.Split(s, " ")
	if len(t) != len(pattern) {
		return false
	}
	cnt1, cnt2 := make(map[string]byte), make(map[byte]string)
	for i, ss := range t {
		c := pattern[i]
		if v, ok := cnt1[ss]; ok && v != c {
			return false
		}
		if v, ok := cnt2[c]; ok && v != ss {
			return false
		}
		cnt1[ss] = c
		cnt2[c] = ss
	}
	return true
}

// func wordPattern(pattern string, s string) bool {
// 	m, n := len(pattern), len(s)
	
// 	cnt1 := [26]string{}
// 	cnt2 := make(map[string]byte)
// 	i, j, k := 0, 0, 0
// 	for i < n && j < n {
// 		i = j
// 		for i < n && s[i] == ' ' {
// 			i++
// 		}
// 		j = i
// 		for j < n && s[j] != ' ' {
// 			j++
// 		}
// 		if i == n {
// 			break
// 		}
// 		x := s[i:j]
		
// 		v1 := cnt1[pattern[k] - 'a'] 
// 		v2, ok := cnt2[x]
// 		if v1 != "" && x != v1 {
// 			return false
// 		}
// 		if ok && v2 != pattern[k] {
// 			return false
// 		}
// 		cnt1[pattern[k] - 'a'] = x
// 		cnt2[x] = pattern[k]
// 		k++
// 	}
// 	return true
// }
