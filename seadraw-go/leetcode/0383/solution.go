package solution

func canConstruct(ransomNote string, magazine string) bool {
	if len(ransomNote) > len(magazine) {
        return false
    }
	freq := [26]int{}
	for _, c :=  range magazine {
		freq[c - 'a']++
	}
	for _, c := range ransomNote {
		freq[c - 'a']--
		if freq[c - 'a'] < 0 {
			return false
		}
	}
	return true
}
