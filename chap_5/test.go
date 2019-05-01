package main

import "fmt"

type User struct {
    Id   int
    Name string
}

func NewUser(id int, name string) *User {
    u := new(User)
    u.Id = id
    u.Name = name
    return u
}

func main() {
    user := NewUser(1, "Bob")
	fmt.Println(user)
    fmt.Printf("id: %d, name: %s\n", user.Id, user.Name)
}
