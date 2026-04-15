package solution

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	t.Run("test01", func(t *testing.T) {
		got := lengthOfLongestSubstring("abba")
		if got != 2 {
			t.Fatalf("lengthOfLongestSubstring(abba) = %d, want 2", got)
		}
	})
}
