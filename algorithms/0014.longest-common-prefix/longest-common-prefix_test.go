package leetcode0014

import (
	"testing"
)

type question struct {
	q []string
	a string
}

func TestLongestCommonPrefix(t *testing.T) {
	showDetail := true
	qss := []question{
		question{[]string{}, ""},
		question{[]string{""}, ""},
		question{[]string{"abcd"}, "abcd"},
		question{[]string{"", ""}, ""},
		question{[]string{"abcd", ""}, ""},
		question{[]string{"aa", "aaaa", "aaaa"}, "aa"},
		question{[]string{"aaaa", "aa", "aaaa"}, "aa"},
		question{[]string{"flow", "flower", "flight"}, "fl"},
		question{[]string{"flower", "flow", "flight"}, "fl"},
	}

	for _, qs := range qss {
		ans := longestCommonPrefix(qs.q)
		if ans == qs.a {
			t.Logf("OK")
		} else {
			t.Logf("FAILED")
		}
		if showDetail {
			t.Logf("  longestCommonPrefix(%v) get \"%s\", expect \"%s\".", qs.q, ans, qs.a)
		}
	}
}
