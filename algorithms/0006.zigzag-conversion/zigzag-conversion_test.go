package leetcode0006

import (
	"testing"
)

func TestConvert(t *testing.T) {
	t.Log(convert("PAYPALISHIRING", 3) == "PAHNAPLSIIGYIR")
	t.Log(convert("PAYPALISHIRING", 4) == "PINALSIGYAHRPI")
	t.Log(convert("PAY", 4) == "PAY")
	t.Log(convert("ABCDEFG", 2) == "ACEGBDF")
	t.Log(convert("ABCDEFG", 1) == "ABCDEFG")
	t.Log(convert("", 2) == "")

}
