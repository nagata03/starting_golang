package main

import "fmt"

type Point struct { X, Y int }

func (p *Point) Render() {
	fmt.Printf("(%d, %d)\n", p.X, p.Y)
}

func main() {
    f := (*Point).Render
	fmt.Printf("%T\n", f)
	f(&Point{X:1, Y:2})
}
