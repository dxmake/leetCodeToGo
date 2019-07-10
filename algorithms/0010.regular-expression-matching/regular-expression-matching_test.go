package leetcode0010

import (
	"testing"
)

type question struct {
	qS, qP string
	a      bool
}

func TestIsMatch(t *testing.T) {
	showDetail := true
	qss := []question{
		question{"", "", true},
		question{"aa", "a.", true},
		question{"aab", "a*.*", true},
		question{"aab", ".*b", true},
		question{"aa", "b*", false},
		question{"aab", "a*", false},
	}

	for _, qs := range qss {
		ans := isMatch(qs.qS, qs.qP)
		if ans == qs.a {
			t.Log("OK")
		} else {
			t.Log("FAILED")
		}
		if showDetail {
			t.Logf("  isMatch(\"%s\", \"%s\") get %v, expect %v.", qs.qS, qs.qP, ans, qs.a)
		}
	}

}
