package main

import "fmt"

func isPalindrome(x int) bool {
	y := x
	tmp := 0
	if x > 0 && x%10 != 0 {
		for y != 0 {
			tmp = tmp*10 + y%10
			y = y / 10
		}

	}

	if tmp == x || x == 0 {
		return true
	}

	return false
}

func main() {
	x := 12321
	res := isPalindrome(x)
	fmt.Println(res)
}
