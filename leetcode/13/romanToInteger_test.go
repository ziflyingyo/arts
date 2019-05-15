package romantointeger

import "testing"

func TestRomanToInt1(t *testing.T) {
	listresult := romanToInt("IX")
	if listresult != 9 {
		t.Error("result is wrong!")
	} else {
		t.Log("result is right")
	}
}

func TestRomanToInt2(t *testing.T) {
	listresult := romanToInt("LVIII")
	if listresult != 58 {
		t.Error("result is wrong!")
	} else {
		t.Log("result is right")
	}
}

func TestRomanToInt3(t *testing.T) {
	listresult := romanToInt("MCMXCIV")
	if listresult != 1994 {
		t.Error("result is wrong!")
	} else {
		t.Log("result is right")
	}
}
