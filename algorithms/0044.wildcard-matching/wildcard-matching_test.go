package leetcode0044

import (
	"testing"
)

type question struct {
	inS  string
	inP  string
	outA bool
}

func TestIsMatch(t *testing.T) {
	qss := []question{
		question{"", "", true},
		question{"", "a", false},
		question{"a", "", false},
		question{"aab", "*", true},
		question{"aab", "**", true},
		question{"aab", "***", true},
		question{"aab", "**?", true},
		question{"aab", "?*?*?", true},
		question{"aab", "**?b", true},
		question{"aab", "a*?*b", true},
		question{"aab", "a**a", false},
		question{"aab", "a??", true},
	}

	for _, qs := range qss {
		ans := isMatch(qs.inS, qs.inP)
		if ans != qs.outA {
			t.Errorf("FAILED: isMatch(\"%v\",\"%v\") get %v, expect %v.", qs.inS, qs.inP, ans, qs.outA)
		} else {
			t.Logf("OK: isMatch(\"%v\",\"%v\") get %v.", qs.inS, qs.inP, ans)
		}
	}
}
