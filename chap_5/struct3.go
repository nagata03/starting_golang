package main

import (
	"fmt"
)

func main() {
	type Base struct {
		Id    int
		Owner string
	}

	type A struct {
		Base // フィールド名を省略した埋め込み構造体
		Name string
		Area string
	}

	type B struct {
		Base   // フィールド名を省略した埋め込み構造体
		Title  string
		Bodies []string
	}

	a := A{
		Base: Base{17, "Taro"},
		Name: "Taro",
		Area: "Tokyo",
	}

	b := B{
		Base:   Base{81, "Hanako"},
		Title:  "no title",
		Bodies: []string{"A", "B"},
	}

	fmt.Println(a.Id)
	fmt.Println(a.Owner)
	fmt.Println(b.Id)
	fmt.Println(b.Owner)
}
