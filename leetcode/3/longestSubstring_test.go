package longestsubstring

import "testing"

func TestLongSubstring1(t *testing.T) {
	if lengthOfLongestSubstring("pwwkew") != 3 {
		t.Error("abc result is wrong!")
	} else {
		t.Log("abc result is right")
	}
}
