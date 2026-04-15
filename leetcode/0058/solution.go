package solution

func lengthOfLastWord(s string) int {
	i := len(s) - 1
	for s[i] == ' ' {
		i--
	}
	
	j := i - 1
	for j >= 0 && s[j] != ' ' {
		j--
	}
	
	return i - j
}

// func lengthOfLastWord(s string) int {
// 	n := len(s)
// 	res := 0
// 	flag := false
// 	for i := n - 1; i >= 0; i-- {
// 		if s[i] != ' ' {
// 			res++
// 			flag = true
// 		} else {
// 			if flag {
// 				break
// 			}
// 		}
// 	}
// 	return res
// }
