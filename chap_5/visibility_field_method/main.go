package main

import (
	"fmt"
	"./foo"
)

func main() {
	t := &foo.T{}
	fmt.Println(t.Method1()) // OK
	fmt.Println(t.Field1)    // OK
	//fmt.Println(t.method2()) // コンパイルエラー
	//fmt.Println(t.filed2)    // コンパイルエラー
}


