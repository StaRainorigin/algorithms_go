package solution

func longestCommonPrefix(strs []string) string {
	s0 := strs[0]
	for i, c := range s0 {
		for _, str := range strs {
			if i == len(str) || str[i] != byte(c) {
				return s0[:i]
			}
		}
	}
	return s0
}

// import "strings"

// func longestCommonPrefix(strs []string) string {
// 	m := len(strs)

// 	if m < 1 {
// 		return ""
// 	} else if m == 1 {
// 		return strs[0]
// 	}

// 	var res strings.Builder

// 	n := len(strs[0])
// 	for _, str := range strs {
// 		n = min(n, len(str))
// 	}

// 	for i := 0; i < n; i++ {
// 		cur := strs[0][i]
// 		for j := 1; j < m; j++ {
// 			if strs[j][i] != cur {
// 				return res.String()
// 			}
// 		}
// 		res.WriteByte(cur)
// 	}

// 	return res.String()
// }
