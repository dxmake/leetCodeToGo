package leetcode0008

import (
	"testing"
)

type question struct {
	q string
	a int
}

func TestMyAtoi(t *testing.T) {
	const showDetail bool = true
	qss := []question{
		question{q: "42", a: 42},
		question{q: "     -42", a: -42},
		question{q: "4193 with words", a: 4193},
		question{q: " words and 987", a: 0},
		question{q: "", a: 0},
		question{q: "-91283472332", a: -2147483648},
		question{q: " +-64", a: 0},
		question{q: "+6", a: 6},
		question{q: "-i", a: 0},
		question{q: "     ", a: 0},
		question{q: "000000", a: 0},
		question{q: "-000zigzag", a: 0},
		// Fuck!
		question{q: "9223372036854775808", a: 2147483647},
	}

	for _, qs := range qss {
		ans := myAtoi(qs.q)
		t.Log(ans == qs.a)
		if showDetail {
			t.Logf("  myAtoi(\"%s\") get %d, and expect %d.", qs.q, ans, qs.a)
		}
	}
}
