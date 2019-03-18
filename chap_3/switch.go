package main

import (
	"fmt"
)

func main() {
	switch1(3)
	switch1(5)

	switch2(2)
	switch2(false)
	switch2(nil)
}

func switch1(n int) {
	switch n {
	case 1, 2:
		fmt.Println("1 or 2")
	case 3, 4:
		fmt.Println("3 or 4")
	default:
		fmt.Println("unknown")
	}
}

func switch2(n interface{}) {
	switch n.(type) {
	case bool:
		fmt.Println("bool")
	case int, uint:
		fmt.Println("integer or unsigned integer")
	case string:
		fmt.Println("string")
	default:
		fmt.Println("unknown type")
	}
}

