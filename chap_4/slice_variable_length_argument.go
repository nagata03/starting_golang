package main

func main() {
	println(sum(1, 2, 3))
	println(sum(1, 2, 3 ,4, 5))
	println(sum())
	s := []int{1, 2, 3}
	println(sum(s...))
}

func sum(s ...int) int {
	n := 0
	for _, v := range s {
		n += v
	}
	return n
}

