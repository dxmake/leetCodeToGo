package leetcode0007

import (
	"testing"
)

type questions struct {
	q int
	a int
}

func TestReverse(t *testing.T) {
	const showDetail bool = true
	qss := []questions{
		questions{q: 0, a: 0},

		questions{q: 1, a: 1},
		questions{q: -1, a: -1},

		questions{q: 120, a: 21},
		questions{q: -120, a: -21},

		questions{q: 123, a: 321},
		questions{q: -123, a: -321},

		questions{q: 2147483647, a: 0},
		questions{q: -2147483647, a: 0},

		questions{q: 2147483648, a: 0},
		questions{q: -2147483648, a: 0},

		questions{q: 7463847412, a: 2147483647},
		questions{q: 8463847412, a: 0},

		questions{q: -8463847412, a: -2147483648},
		questions{q: -9463847412, a: 0},

		questions{q: 2000000001, a: 1000000002},
		questions{q: -2000000001, a: -1000000002},

		questions{q: 12345678901, a: 0},
		questions{q: -12345678901, a: 0},
	}

	for _, qs := range qss {
		ans := reverse(qs.q)
		t.Log(ans == qs.a)
		if showDetail {
			t.Logf("  reverse(%d) get %d, and expect %d", qs.q, ans, qs.a)
		}
	}
}
