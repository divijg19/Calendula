package main

import (
	"fmt"
	"math"
	"math/rand"
)

var (
	c, python, java bool
	i, j                       = 1, 2
	ToBe            bool       = false
	MaxInt          uint64     = 1<<64 - 1
	z               complex128 = complex(1, 2)
)

func add(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println("My favourite number is", rand.Intn(10))
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
	fmt.Println(math.Pi)
	fmt.Println(add(42, 13))
	a, b := swap("hello", "world")
	fmt.Println(a, b)
	fmt.Println(split(17))
	k := 3
	c, python, java := true, false, "no!" //redeclaration inside func
	fmt.Println(i, j, k, c, python, java)
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
	var i int
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	var d bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, d, s)
	fmt.Println(x, y, z)
}
