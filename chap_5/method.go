package main

import (
	"fmt"
)

type IntPair [2]int

/* ペアの先頭を返すメソッド */
func (ip IntPair) First() int {
	return ip[0]
}

/* ペアの末尾を返すメソッド */
func (ip IntPair) Last() int {
	return ip[len(ip)-1]
}

type Strings []string

/* 文字列のスライスを区切り文字で連結するメソッド */
func (s Strings) Join(d string) string {
	sum := ""
	for _, v := range s {
		if sum != "" {
			sum += d
		}
		sum += v
	}
	return sum
}

func main() {
	ip := IntPair{1, 2}
	fmt.Println(ip.First)
	fmt.Println(ip.Last)

	fmt.Println(Strings{"A", "B", "C"}.Join("/"))
}
