package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, Golang!!")
	fmt.Println("My", "name", "is", "Taro")

	fmt.Printf("数値=%d\n", 5) // %d
	fmt.Printf("10進数=%d 2進数=%b 8進数=%o 16進数=%x\n", 17, 17, 17, 17)

	fmt.Printf("数値=%v, 文字列=%v, 配列=%v\n", 10, "Golang", [...]int{1, 2, 3})
	fmt.Printf("数値=%#v, 文字列=%#v, 配列=%#v\n", 10, "Golang", [...]int{1, 2, 3})
	fmt.Printf("数値=%T, 文字列=%T, 配列=%T\n", 10, "Golang", [...]int{1, 2, 3})

	print("Hello, World!!")
	println("Hello, Golang!!")
	println(1, 2, 3)
}
