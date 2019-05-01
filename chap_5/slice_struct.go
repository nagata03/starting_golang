package main

import (
	"fmt"
)

func main() {
	type Point struct{ X, Y int }

	ps := make([]Point, 5)
	for _, p := range ps {
		fmt.Println(p.X, p.Y)
	}
}
