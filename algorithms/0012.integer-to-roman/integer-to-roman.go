package leetcode0012

func intToRoman(num int) string {
	mp := [][]rune{
		[]rune("IIIVIIIX"),
		[]rune("XXXLXXXC"),
		[]rune("CCCDCCCM"),
		[]rune("MMM"),
	}
	a := [4]int{}
	a[3], num = num/1000, num%1000
	a[2], num = num/100, num%100
	a[1], num = num/10, num%10
	a[0] = num
	res := ""
	for i := 3; i >= 0; i-- {
		switch a[i] {
		case 0, 1, 2, 3:
			res += string(mp[i][:a[i]])
		case 4:
			res += string(mp[i][2:4])
		case 5, 6, 7, 8:
			res += string(mp[i][3 : a[i]-1])
		case 9:
			res += string(mp[i][6:8])
		}
	}
	return res
}
