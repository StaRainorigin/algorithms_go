// 给定一个字符串 s 和一个字符串数组 words。 words 中所有字符串 长度相同。
//  s 中的 串联子串 是指一个包含  words 中所有字符串以任意顺序排列连接起来的子串。
// 例如，如果 words = ["ab","cd","ef"]， 那么 "abcdef"， "abefcd"，"cdabef"， "cdefab"，"efabcd"， 和 "efcdab" 都是串联子串。 "acdbef" 不是串联子串，因为他不是任何 words 排列的连接。
// 返回所有串联子串在 s 中的开始索引。你可以以 任意顺序 返回答案。

package solution

func findSubstring(s string, words []string) []int {
	m := len(words)
	if m <= 0 {
		return []int{}
	}
	step, n := len(words[0]), len(s)
	if step == 0 {
		return []int{}
	}

	// 统计每个单词应该出现的次数
	wordCount := make(map[string]int)
	for _, word := range words {
		wordCount[word]++
	}

	res := make([]int, 0)

	// 由于所有单词长度相同为 step，我们只需要检查 step 个不同的偏移量
	// 例如 step=3，偏移量 0,1,2 就能覆盖所有可能的起始位置
	for offset := range step {
		// 记录当前窗口中每个单词出现的次数
		records := make(map[string]int)
		count := 0

		// 滑动窗口：以 step 为步长遍历
		for i := offset; i+step <= n; i += step {
			// 获取当前位置的单词
			cur := s[i : i+step]

			// 如果是目标单词，加入窗口
			if wordCount[cur] > 0 {
				records[cur]++
				count++
			}

			// 当窗口大小超过 m*step 时，移除最左边的单词
			if i >= m*step {
				pre := s[i-m*step : i-m*step+step]
				if wordCount[pre] > 0 {
					records[pre]--
					count--
				}
			}

			// 窗口大小正好为 m 个单词时，检查是否匹配
			if count == m {
				flag := true
				for word, cnt := range wordCount {
					// 比较当前窗口中单词出现次数与目标是否一致
					if records[word] != cnt {
						flag = false
						break
					}
				}
				// 完全匹配，记录起始索引
				if flag {
					// 起始位置 = 窗口右边界 - 窗口总长度 + step
					res = append(res, i-(m-1)*step)
				}
			}
		}
	}
	return res
}
