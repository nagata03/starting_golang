package main

import (
	"fmt"
)

func main() {
	f := later()

	fmt.Println(f("Golang"))
	fmt.Println(f("is"))
	fmt.Println(f("fun!!"))

	ints := integers()
	fmt.Println(ints())
	fmt.Println(ints())
	fmt.Println(ints())
}

func later() func(string) string {
	// 1つ前に与えられた文字列を保存するための変数
	var store string

	// 引数に文字列を取り文字列を返す関数を定義
	return func(next string) string {
		s := store
		store = next
		return s
	}
}

func integers() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

