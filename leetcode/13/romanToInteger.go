package romantointeger

func romanToInt(s string) int {
	var result int
	dict := map[string]int{"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000}
	for i := 0; i < len(s); i++ {
		if i < len(s)-1 && dict[string(s[i])] < dict[string(s[i+1])] {
			result = result - dict[string(s[i])]
		} else {
			result = result + dict[string(s[i])]
		}
	}

	return result
}
