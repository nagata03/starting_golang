package main

import (
	"fmt"
)

func main() {
	type (
		MyInt       int
		IntPair     [2]int
		Strings     []string
		AreaMap     map[string][2]float64
		IntsChannel chan []int
	)

	var n1 MyInt = 5
	n2 := MyInt(7)
	fmt.Println(n1)
	fmt.Println(n2)

	pair := IntPair{1, 2}
	strs := Strings{"Apple", "Banana", "Cherry"}
	amap := AreaMap{"Tokyo", {35.689488, 139.691706}}
	ich := make(IntsChannel)

}
