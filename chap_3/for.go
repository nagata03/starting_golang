package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 10; i++ {
		if (i % 2 == 1) {
			continue
		}
		fmt.Println(i)
	}

	fruits := [3]string{"Apple", "Banana", "Cherry"}
	for i, s := range fruits {
		fmt.Printf("fruits[%d]=%s\n", i, s)
	}

	for i, s := range "ABC" {
		fmt.Printf("[%d] -> %d\n", i, s)
	}
}

