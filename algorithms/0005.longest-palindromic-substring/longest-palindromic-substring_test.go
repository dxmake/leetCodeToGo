package leetcode0005

import (
	"testing"
)

func TestLongestPalindrome(t *testing.T) {
	t.Log(longestPalindrome("cbbd") == "bb")
	t.Log(longestPalindrome("") == "")
	t.Log(longestPalindrome("ab") == "a")
	t.Log(longestPalindrome("abcbaab") == "abcba")
	t.Log(longestPalindrome("babad") == "bab")
}
