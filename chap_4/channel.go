package main

import (
	"fmt"
)

func main() {
	ch := make(chan int, 10) // 第2引数はバッファサイズ
	ch <- 5 // チャネルに整数5を送信
	i := <-ch // チャネルから整数値を受信）
	fmt.Println(i)

	fmt.Println(len(ch)) // lenで取得できるのはチャネルのバッファ内に溜められているデータの個数
	ch <- 10
	ch <- 100
	fmt.Println(len(ch))

	close(ch)
	i, ok := <-ch
	fmt.Printf("%d, %t\n", i, ok)
}
