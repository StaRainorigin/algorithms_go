package solution

import "math"

type TrieNode struct {
	ch         [26]*TrieNode
	minNodeLen int
	minNodeIdx int
}

func newTrie(ss []string) *TrieNode {
	root := &TrieNode{minNodeLen: math.MaxInt}
	for i, s := range ss {
		// 更新 root
		if len(s) < root.minNodeLen {
			root.minNodeLen = len(s)
			root.minNodeIdx = i
		}
		cur := root
		for j := len(s) - 1; j >= 0; j-- {
			c := s[j] - 'a'
			if cur.ch[c] == nil {
				cur.ch[c] = &TrieNode{minNodeLen: math.MaxInt}
			}
			cur = cur.ch[c]
			if len(s) < cur.minNodeLen {
				cur.minNodeLen = len(s)
				cur.minNodeIdx = i
			}
		}
	}
	return root
}

func stringIndices(wordsContainer []string, wordsQuery []string) []int {
	trieTree := newTrie(wordsContainer)
	
	ans := make([]int, len(wordsQuery))
	for i, q := range wordsQuery {
		cur := trieTree
		for j := len(q) - 1; j >= 0; j-- {
			c := q[j] - 'a'
			if cur.ch[c] == nil {
				break
			} 
			cur = cur.ch[c]
		}
		ans[i] = cur.minNodeIdx
	}

	return ans
}
