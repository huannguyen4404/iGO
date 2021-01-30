package main

import (
	"fmt"
	"math"
)

// func add(x int, y int) int {
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

//var c, python, java bool

//var i, j int = 1, 2

//var (
//	ToBe   bool       = false
//	MaxInt uint64     = 1<<64 - 1
//	z      complex128 = cmplx.Sqrt(-5 + 12i)
//)
const Pi = 3.14

func main() {
	//fmt.Println("My favorite number is", rand.Intn(10))
	//fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
	//fmt.Println(math.Pi)
	//fmt.Println(add(42, 13))

	//a, b := swap("hello", "world")
	//fmt.Println(a, b)
	//
	//fmt.Println(split(17))

	//var i, j int = 1, 2
	//k := 3
	//c, python, java := true, false, "no!"
	//
	//fmt.Println(i, j, k, c, python, java)
	//
	//fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	//fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	//fmt.Printf("Type: %T Value: %v\n", z, z)

	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)

	var x, y int = 3, 4
	f = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)

	const World = "Henry"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}
