package solution

import "strings"

func reverseWords(s string) string {
	var res strings.Builder
	j := len(s) - 1
	for j >= 0 {
		for j >= 0 && s[j] == ' ' {
			j--
		}
		if j < 0 {
			break
		}
		
		i := j
		for i >= 0 && s[i] != ' ' {
			i--
		}
		
		if res.Len() > 0 {
			res.WriteByte(' ')
		}
		for k := i + 1; k <= j; k++ {
			res.WriteByte(s[k])
		}
		
		j = i - 1
	}
	return res.String()
}

// import "strings"

// func reverseWords(s string) string {
// 	words := strings.Fields(s)
// 	var builder strings.Builder
	
// 	for i := len(words) - 1; i >= 0; i-- {
// 		builder.WriteString(words[i])
// 		if i > 0 {
// 			builder.WriteString(" ")
// 		}
// 	}
	
// 	return builder.String()
// }