package main

import (
	"fmt"
)

/* 文字列かできることを示すインタフェース */
type Stringfy interface {
	ToString() string
}

/* 構造体型Personの定義 */
type Person struct {
	Name string
	Age  int
}

func (p *Person) ToString() string {
	return fmt.Sprintf("%s(%d)", p.Name, p.Age)
}

/* 構造体型Carの定義 */
type Car struct {
	Number string
	Model  string
}

func (c *Car) ToString() string {
	return fmt.Sprintf("[%s] %s", c.Number, c.Model)
}

func main() {
	/* 異なる型を共通するインタフェース型にまとめることができる */
	vs := []Stringfy{
		&Person{Name: "Taro", Age: 21},
		&Car{Number: "XXX-0123", Model: "PX512"},
	}
	for _, v := range vs {
		fmt.Println(v.ToString())
	}
}
