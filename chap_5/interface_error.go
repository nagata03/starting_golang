package main

import (
	"fmt"
)

type MyError struct {
	Message string
	ErrCode int
}

/* errorインタフェースのメソッドを実装 */
func (e MyError) Error() string {
	return e.Message
}

func RaiseError() error {
	return MyError{Message: "エラーが発生しました", ErrCode: 1234}
}

func main() {
	err := RaiseError()
	fmt.Println(err.Error())
}
