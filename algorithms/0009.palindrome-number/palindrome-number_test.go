package leetcode0009

import (
	"testing"
)

type question struct {
	q int
	a bool
}

func TestIsPalindrome(t *testing.T) {
	var showDetail bool = true
	qss := []question{
		question{q: 121, a: true},
		question{q: -121, a: false},
		question{q: 10, a: false},
		question{q: 0, a: true},
	}

	for _, qs := range qss {
		ans := isPalindrome(qs.q)
		t.Log(ans == qs.a)
		if showDetail {
			t.Logf("  isPalindorme(%d) get %v, expect %v.", qs.q, ans, qs.a)
		}
	}
}
