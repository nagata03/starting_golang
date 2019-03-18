package main

import (
	"./foo"
)

func main() {
	println(foo.MAX)
	//println(foo.internal_const) コンパイルエラー
	println(foo.FooFunc(5))
	//println(foo.internalFunc(10)) コンパイルエラー
}
