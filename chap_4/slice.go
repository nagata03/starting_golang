package main

import (
	"fmt"
)

func main() {
	s := make([]int, 10) // []int型のスライスを生成
	fmt.Println(s)

	a := make([]float64, 3)
	fmt.Println(a)
	a[0] = 3.14
	fmt.Println(a)
	fmt.Println(a[0])
	//fmt.Println(a[3]) // ランタイムパニック発生

	b := make([]int, 8)
	fmt.Println(len(b))
	fmt.Println(cap(b))

	c := make([]int, 5, 10)
	fmt.Println(len(c))
	fmt.Println(cap(c))

	d := []int{1, 2, 3, 4, 5}
	fmt.Println(len(d))
	fmt.Println(cap(d))

	e := d[0:2]
	fmt.Println(e)
	fmt.Println(cap(e))

	f := []int{1, 2, 3}
	f = append(f, 4)
	fmt.Println(f)
	f = append(f, 5, 6, 7)
	fmt.Println(f)

	s0 := []int{1, 2, 3}
	s1 := []int{4, 5, 6}
	s2 := append(s0, s1...)
	fmt.Println(s2)

	/* (A) */
	t := make([]int, 0, 0)
	fmt.Printf("(A) len=%d, cap=%d\n", len(t), cap(t))
	/* (B) */
	t = append(t, 1)
	fmt.Printf("(B) len=%d, cap=%d\n", len(t), cap(t))
	/* (C) */
	t = append(t, []int{2, 3, 4}...)
	fmt.Printf("(C) len=%d, cap=%d\n", len(t), cap(t))
	/* (D) */
	t = append(t, 5)
	fmt.Printf("(D) len=%d, cap=%d\n", len(t), cap(t))
	/* (E) */
	t = append(t, 6, 7, 8, 9)
	fmt.Printf("(E) len=%d, cap=%d\n", len(t), cap(t))

	u0 := []int{1, 2, 3, 4, 5}
	u1 := []int{10, 11}
	n  := copy(u0, u1)
	fmt.Println(n)
	fmt.Println(u0)

	g  := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	g1 := g[2:4]
	fmt.Println(g1)
	println(len(g1))
	println(cap(g1))
	g2 := g[2:4:6]
	fmt.Println(g2)
	println(len(g2))
	println(cap(g2))
}


