package main

import (
	"fmt"
)

var plusAlias = plus
var divAlias  = div

func main() {
	println(plus(1, 2))

	hello()

	q, r := div(10, 2)
	fmt.Printf("商=%d, 余=%d\n", q, r)

	// 無名関数
	f := func(x, y int) int { return x + y }
	println(f(3, 4))
	fmt.Printf("%T\n", f)

	println(plusAlias(100, 50))
	qq, rr := divAlias(100, 25)
	fmt.Printf("商=%d, 余=%d\n", qq, rr)

	ff := returnFunc()
	ff()
	returnFunc()()

	callFunction(func() {
		fmt.Println("I'm a Function.")
	})
}

func plus(x, y int) int {
	return x + y
}

func hello() {
	fmt.Println("Hello, World!!")
	return
}

func div(a, b int) (int, int) {
	q := a / b
	r := a % b
	return q, r
}

func returnFunc() func() {
	return func() {
		fmt.Println("I'm a function.")
	}
}

func callFunction(f func()) {
	f()
}

