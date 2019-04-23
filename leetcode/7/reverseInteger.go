package main

import "fmt"

func reverse(x int) int {
	y := int32(x)
	var res int32
	for y != 0 {
		temp := res*10 + y%10
		if temp/10 != res {
			return 0
		}
		res = temp
		y = y / 10
	}
	return int(res)

}

func main() {
	x := -12000000000000000
	y := reverse(x)
	fmt.Println(y)
}
