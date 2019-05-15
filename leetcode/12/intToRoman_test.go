package inttoroman

import "testing"

func TestIntToRoman1(t *testing.T) {
	if intToRoman(1994) != "MCMXCIV" {
		t.Error("result is wrong!")
	} else {
		t.Log("result is right")
	}

}
