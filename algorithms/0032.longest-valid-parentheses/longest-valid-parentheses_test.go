package leetcode0032

import (
	"testing"
)

type question struct {
	q string
	a int
}

func TestLongestValidParentheses(t *testing.T) {
	showDetail := true
	qss := []question{
		question{"", 0},
		question{"()", 2},
		question{")(", 0},
		question{")()()(", 4},
		question{")()())", 4},
		question{")((()())())(", 10},
		question{")((())", 4},
		question{")()))))", 2},
		question{")())(()))", 4},
		question{"()(()", 2},
	}

	for _, qs := range qss {
		ans := longestValidParentheses(qs.q)
		if ans == qs.a {
			t.Log("OK")
		} else {
			t.Error("FAILED")
		}
		if showDetail {
			t.Logf("  longestValidParentheses(\"%s\") get %d, expect %d.", qs.q, ans, qs.a)
		}
	}
}
