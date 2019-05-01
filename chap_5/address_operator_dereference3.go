package main

import (
	"fmt"
)

func main() {
	p := &[3]int{1, 2, 3}
	pow(p)
	fmt.Println(p)
}

func pow(p *[3]int) {
	i := 0
	for i < 3 {
		(*p)[i] = (*p)[i] * (*p)[i]
		i++
	}
}
