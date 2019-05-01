package main

import (
	"fmt"
)

func main() {
	p := &[3]string{"Apple", "Banana", "Cherry"}

	for i, v := range p {
		fmt.Println(i, v)
	}
}
