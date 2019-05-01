package main

import (
	"fmt"
)

func main() {
	i := 5
	ip := &i
	fmt.Printf("type=%T, address=%p\n", ip, ip)
}
