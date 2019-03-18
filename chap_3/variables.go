package main

var k = 100 // パッケージ変数

func main() {
	var i int
	var x, y, z int
	var (
		a, b int
		name string
	)
	i = 5
	x, y, z = 1, 2, 3
	name = "Golang"
	println("i=", i)
	println("x=", x, "y=", y, "z=", z)
	println("a=", a, "b=", b, "name=", name)

	j := 10
	println("j=", j)

	n := one()
	println("n=", n)

	k = k + 1
	println("k=", k)
}

func one() int {
	return 1
}
