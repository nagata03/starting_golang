package main

import (
	"fmt"
)

func main() {
	LOOP:
		for i := 1; i <= 3; i++ {
			for j := 1; j <= 3; j++ {
				if j > 2 {
					continue LOOP
				}
				fmt.Printf("%d * %d = %d\n", i, j, i * j)
			}
			fmt.Println("ここには来ない")
		}
}

