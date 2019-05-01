package main

import (
	"fmt"
)

func main() {
	type Feed struct {
		Name   string
		Amount uint
	}

	type Animal struct {
		Name string
		Feed Feed
	}

	a := Animal{
		Name: "Monkey",
		Feed: Feed{
			Name:   "Banana",
			Amount: 10,
		},
	}

	fmt.Println(a)
	fmt.Println(a.Name)
	fmt.Println(a.Feed.Name)
}
