package main

import "fmt"

/* Stringfyインタフェースを定義 */
type Stringfy interface {
    ToString() string
}

/* 構造体型Userを定義 */
type User struct {
    Id   int
    Name string
}

/* *User型にToStringメソッドを定義 */
func (u *User) ToString() string {
    return fmt.Sprintf("ID: %d, Name: %s", u.Id, u.Name)
}

/* 構造体型Bookを定義 */
type Book struct {
    Id    int
    Title string
}

/* *Book型にToStringメソッドを定義 */
func (b *Book) ToString() string {
    return fmt.Sprintf("[ID]%d, [Title]%s", b.Id, b.Title)
}

func Println(s Stringfy) {
    fmt.Println(s.ToString())
}

func main() {
	Println(&User{Id: 1, Name: "Jack"})
	Println(&Book{Id: 101, Title: "スターティングGo言語"})
}
