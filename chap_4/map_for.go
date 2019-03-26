package main

import (
	"fmt"
)

func main() {
	m := map[int]string{
		1: "Apple",
		2: "Banana",
		3: "Cherry",
	}

	for k, v := range m {
		fmt.Printf("%d => %s\n", k, v)
	}

	fmt.Println(len(m))
}
