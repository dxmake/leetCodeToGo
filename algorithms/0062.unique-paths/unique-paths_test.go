package leetcode0062

import (
	"testing"
)

type question struct {
	qm int
	qn int
	a  int
}

func TestUniquePaths(t *testing.T) {
	qss := []question{
		question{1, 1, 1},
		question{1, 10, 1},
		question{10, 1, 1},
		question{2, 2, 2},
		question{2, 5, 5},
		question{5, 2, 5},
		question{3, 3, 6},
		question{3, 7, 28},
		question{7, 3, 28},
		question{7, 7, 924},
	}
	for _, qs := range qss {
		ans := uniquePaths(qs.qm, qs.qn)
		if ans != qs.a {
			t.Errorf("FAILED: uniquePaths(%d, %d) get %d, expect %d.", qs.qm, qs.qn, ans, qs.a)
		} else {
			t.Logf("OK: uniquePaths(%d, %d) get %d.", qs.qm, qs.qn, ans)
		}
	}
}
