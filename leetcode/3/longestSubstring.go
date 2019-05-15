package longestsubstring

import "strings"

func lengthOfLongestSubstring(s string) int {
	max := 0
	for i := 0; i < len(s); i++ {
		for j := i + 1; j <= len(s); j++ {
			if !strings.Contains(s[i:j-1], s[j-1:j]) {
				if max < j-i {
					max = j - i
				}
			} else {
				break
			}
		}

	}
	return max

}
