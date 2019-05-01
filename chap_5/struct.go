package main

import (
	"fmt"
)

func main() {
	type Point struct {
		X, Y int
	}

	var pt Point
	fmt.Println(pt.X)
	fmt.Println(pt.Y)

	pt.X = 10
	pt.Y = 8
	fmt.Println(pt.X)
	fmt.Println(pt.Y)

	pt2 := Point{X: 1, Y: 2}
	fmt.Println(pt2.X)
	fmt.Println(pt2.Y)

	pt3 := Point{Y: 28}
	fmt.Println(pt3.X)
	fmt.Println(pt3.Y)
	fmt.Println(pt3)
}
