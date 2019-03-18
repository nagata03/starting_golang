package main

import (
	"fmt"
)

func main() {
	assertion(1000)
	assertion(nil)
	assertion("hoge")
	assertion(3.14)
}

func assertion(x interface{}) {
	if x == nil {
		fmt.Println("x is nil")
	} else if i, isInt := x.(int); isInt {
		fmt.Printf("x is integer : %d\n", i)
	} else if s, isString := x.(string); isString {
		fmt.Printf("x is string: %s\n", s)
	} else {
		fmt.Println("unsupported type!!")
	}
}

