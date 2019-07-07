package leetcode0008

func myAtoi(str string) int {
	// 状态机的状态
	const (
		stSpace = iota // 空格
		stSign         // 符号
		stNum          // 数字
		stEnd          // 数字结束
		stFail         // 异常
		numSt
	)
	// Ascii 字符定义
	const (
		typSpace = iota
		typSign
		typNum
		typOther
		typNil
		numTyp
	)
	// Ascii 字符类型查询表
	typAscii := [128]int{}
	for i := 0; i < 128; i++ {
		switch {
		case i == 0:
			typAscii[i] = typNil
		case rune(i) == ' ':
			typAscii[i] = typSpace
		case rune(i) == '+' || rune(i) == '-':
			typAscii[i] = typSign
		case i >= 48 && i <= 57:
			typAscii[i] = typNum
		default:
			typAscii[i] = typOther
		}
	}
	// 状态机转移
	stm := [numSt][numTyp]int{
		{stSpace, stSign, stNum, stFail, stFail},
		{stFail, stFail, stNum, stFail, stFail},
		{stEnd, stEnd, stNum, stEnd, stEnd},
		{stEnd, stEnd, stEnd, stEnd, stEnd},
		{stFail, stFail, stFail, stFail, stFail},
	}
	// 最大最小值
	const minInt int = -2147483648
	const maxInt int = 2147483647

	state := stSpace
	res := 0
	signNeg := false
	toEnd := false
	var nextCh rune

	for i := -1; i <= len(str); i++ {
		switch state {
		case stSpace:
		case stSign:
			switch str[i] {
			case '+':
				signNeg = false
			case '-':
				signNeg = true
			}
		case stNum:
			res = 10*res + int(str[i]) - 48
			if res > maxInt {
				toEnd = true
			}
		case stEnd:
			if signNeg {
				res = -res
				signNeg = false
			}
			switch {
			case res > maxInt:
				res = maxInt
			case res < minInt:
				res = minInt
			}
			return res
		case stFail:
			return 0
		}
		if i+1 == len(str) {
			nextCh = rune(0)
		} else {
			nextCh = rune(str[i+1])
		}
		state = stm[state][typAscii[nextCh]]
		if toEnd {
			state = stEnd
		}
	}
	// won't reach this point
	return res
}
