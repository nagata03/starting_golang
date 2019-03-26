package main

import (
	"fmt"
)

func main() {
	m := make(map[int]string)

	m[1]  = "US"
	m[81] = "Japan"
	m[86] = "China"

	fmt.Println(m)

	m[86] = "India"
	fmt.Println(m)

	m1 := map[int]string{1: "Taro", 2: "Hanako", 3: "Jiro"}
	fmt.Println(m1)

	m2 := map[int][]int{
		1: []int{1},
		2: []int{1, 2},
		3: {1, 2, 3}, // 内側のスライスリテラルの[]intは省略可能
	}
	fmt.Println(m2)

	m3 := map[int]map[float64]string{
		1: {3.14: "円周率"},
	}
	fmt.Println(m3)

	m4 := map[int]string{1: "A", 2: "B", 3: "C"}
	fmt.Println(m4[1])
	fmt.Println(m4[4])
	s, ok := m4[1]
	fmt.Printf("%s, %t\n", s, ok)
	s1, ok1 := m4[4]
	fmt.Printf("%s, %t\n", s1, ok1)
}

