package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("uint32 max value = %d\n", math.MaxUint32)
	doSomething(1, 3)

	s := `
Goの
RAW文字列リテラルによる
複数行に渡る
文字列
`
	fmt.Printf("%v", s)

	a := [3]int{1, 2}
	b := [...]string{"hoge", "fuga", "foo", "bar"}
	fmt.Printf("%v\n", a)
	fmt.Printf("%v\n", b)

	var x interface{}
	fmt.Printf("%v\n", x)
}

func doSomething(a, b uint32) bool {
	if (math.MaxUint32 - a) < b {
		// オーバーフローするのでfalse
		return false
	} else {
		// チェック済みのため問題なし
		sum := a + b
		println("sum=", sum)
		return true
	}
}
