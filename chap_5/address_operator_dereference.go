package main

import (
	"fmt"
)

func main() {
	var i int

	p := &i // 任意の型からポインタ型を生成
	fmt.Printf("%T\n", p)

	pp := &p
	fmt.Printf("%T\n", pp)

	i = 5
	fmt.Println(*p)
	*p = 10
	fmt.Println(i)
}
