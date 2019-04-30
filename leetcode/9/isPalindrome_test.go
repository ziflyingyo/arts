package main

import "testing"

func TestIsPalindrome1(t *testing.T) {
	if isPalindrome(1233321) != true {
		t.Error("1233321 result is wrong!")
	} else {
		t.Log("1233321 result is right")
	}
}

func TestIsPalindrome2(t *testing.T) {
	if isPalindrome(1233320) == true {
		t.Error("1233320 result is wrong!")
	} else {
		t.Log("1233320 result is right")
	}
}

func TestIsPalindrome3(t *testing.T) {
	if isPalindrome(-1233321) == true {
		t.Error("-1233321 result is wrong!")
	} else {
		t.Log("-1233321 result is right")
	}
}

func TestIsPalindrome4(t *testing.T) {
	if isPalindrome(0) != true {
		t.Error("0 result is wrong!")
	} else {
		t.Log("0 result is right")
	}
}
