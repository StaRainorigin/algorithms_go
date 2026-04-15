package solution

// strStrKMP 返回 needle 在 haystack 中首次出现的位置；不存在返回 -1
func strStrKMP(haystack, needle string) int {
	if len(needle) == 0 {
		return 0
	}

	lps := buildLPS(needle) // longest prefix suffix 数组
	j := 0                  // needle 指针

	for i := 0; i < len(haystack); i++ {
		// 失配时，根据 lps 回退 j，而不是回退 i
		for j > 0 && haystack[i] != needle[j] {
			j = lps[j-1]
		}

		// 匹配则推进 j
		if haystack[i] == needle[j] {
			j++
		}

		// j 走到 needle 末尾，说明找到一次完整匹配
		if j == len(needle) {
			return i - len(needle) + 1
		}
	}

	return -1
}

// buildLPS 构建 lps 数组：lps[i] 表示 needle[:i+1] 的最长相等前后缀长度
func buildLPS(pattern string) []int {
	lps := make([]int, len(pattern))
	length := 0 // 当前最长前后缀长度

	for i := 1; i < len(pattern); i++ {
		for length > 0 && pattern[i] != pattern[length] {
			length = lps[length-1]
		}
		if pattern[i] == pattern[length] {
			length++
		}
		lps[i] = length
	}

	return lps
}