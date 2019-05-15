package main

import "fmt"

func trace(s string) string {
	fmt.Println("entering:", s)
	return s
}

func un(s string) {
	fmt.Println("leaving:", s)
}

func a() {
	defer un(trace("a"))
	fmt.Println("in a")
}

func b() {
	defer un(trace("b"))
	fmt.Println("in b")
	a()
}

func main() {
	b()
}

/*
result is:
entering: b
in b
entering: a
in a
leaving: a
leaving: b

因为defer声明时会先计算确定参数的值，defer推迟执行的只是其函数体。
*/
