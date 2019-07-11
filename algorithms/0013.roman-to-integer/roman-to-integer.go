package leetcode0013

func romanToInt(s string) int {
	mp := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	sum := 0
	tmpSum := 0
	for _, c := range s {
		switch {
		case mp[c] < tmpSum:
			sum += tmpSum
			tmpSum = mp[c]
		case mp[c] > tmpSum:
			tmpSum = mp[c] - tmpSum
		case mp[c] == tmpSum:
			tmpSum += mp[c]
		}
	}
	sum += tmpSum
	return sum
}
