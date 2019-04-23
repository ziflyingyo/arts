package main

import "testing"

func TestReverseInteger1(t *testing.T) {
	if reverse(120) != 21 {
		t.Error("120 result is wrong!")
	} else {
		t.Log("120 result is right")
	}

}

func TestReverseInteger2(t *testing.T) {
	if reverse(-123) != -321 {
		t.Error("-123 result is wrong!")
	} else {
		t.Log("-123 result is right")
	}

}
