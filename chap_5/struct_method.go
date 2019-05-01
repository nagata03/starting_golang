package main

import (
	"fmt"
	"math"
)

type Point struct{ X, Y int }

/* *Point型のメソッドRender */
func (p *Point) Render() {
	fmt.Printf("<%d,%d>\n", p.X, p.Y)
}

/* 2点間の距離を求めるメソッド */
func (p *Point) Distance(dp *Point) float64 {
	x, y := p.X-dp.X, p.Y-dp.Y
	return math.Sqrt(float64(x*x + y*y))
}

func main() {
	p := &Point{X: 5, Y: 12}
	p.Render()

	q := &Point{X: 0, Y: 0}
	fmt.Println(p.Distance(q))
}
