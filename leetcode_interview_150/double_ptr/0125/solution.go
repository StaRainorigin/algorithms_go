package solution

func isLetter(letter byte) bool { 
	if ('a' <= letter && letter <= 'z') || ('A' <= letter && letter <= 'Z') || ('0' <= letter && letter <= '9') {		
		return true
	}
	return false
}

func toLower(letter byte) byte {
	if 'A' <= letter && letter <= 'Z' {
		return letter - 'A' + 'a'
	}
	return letter
}

func isPalindrome(s string) bool {
	i, j := 0, len(s) - 1
	for i < j {
		for i < j && !isLetter(s[i]) {
			i++
		}
		for i < j && !isLetter(s[j]) {
			j--
		}
		if toLower(s[i]) == toLower(s[j]) {
			i++
			j--
		} else {
			return false
		}
	}
	return true
}
