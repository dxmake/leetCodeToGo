package leetcode0003

import (
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	t.Log(lengthOfLongestSubstring("abcabcbb") == 3)
	t.Log(lengthOfLongestSubstring("aaaaaa") == 1)
	t.Log(lengthOfLongestSubstring("abababababa") == 2)
	t.Log(lengthOfLongestSubstring("pwwkew") == 3)
	t.Log(lengthOfLongestSubstring("a") == 1)
	t.Log(lengthOfLongestSubstring("") == 0)

}
