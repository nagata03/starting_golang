package main

import (
	"fmt"
)

func main() {
	i := 1
	inc(&i)
	inc(&i)
	inc(&i)
	fmt.Println(i)
}

func inc(p *int) {
	*p++
}
