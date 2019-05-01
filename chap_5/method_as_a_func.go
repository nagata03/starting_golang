package main

import (
	"fmt"
)

type Point struct{ X, Y int }

func (p *Point) ToString() string {
	return fmt.Sprintf("[%d,%d]", p.X, p.Y)
}

func main() {
	/* 変数fは(func(*Point) string)型 */
	f := (*Point).ToString
	fmt.Println(f(&Point{X: 7, Y: 11}))
}
