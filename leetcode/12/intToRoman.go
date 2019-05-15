package inttoroman

func intToRoman(num int) string {
	vals := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	strs := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	var res string
	for i := range vals {
		for num >= vals[i] {
			res = res + strs[i]
			num = num - vals[i]
		}
	}
	return res

}
