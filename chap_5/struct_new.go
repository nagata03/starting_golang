package main

import (
	"fmt"
)

func main() {
	type Person struct {
		Id   int
		Name string
		Area string
	}

	p := new(Person)
	fmt.Printf("%d, %T, %T\n", p.Id, p.Name, p.Area)

	q := &Person{Id: 1, Name: "hoge", Area: "Tokyo"}
	fmt.Printf("%d, %T, %T\n", q.Id, q.Name, q.Area)
}
