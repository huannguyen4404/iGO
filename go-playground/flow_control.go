package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim { // if with a short statement
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)

	}
	// v is not available any more
	return lim
}

func printOS() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {  // switch in Go is not need break,
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Printf("%s.\n", os)
	}
}

func printInitTxt() {
	t := time.Now()
	switch {  // switch without condition. Useful to write long if-then-else chains
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}

func main() {
	defer printOS()  // defer
	printInitTxt()

	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)

	defer fmt.Println(sqrt(2), sqrt(-4))  // stacking defer: call stack: Last in first out.

	sum := 1
	//for i := 0; i < 10; i ++ {
	//for ; sum < 10; {  // init and post statement is optional
	for sum < 10 { // same as while in other languages
		sum += sum
	}
	//for {  // For without loop condition: forever
	//}
	fmt.Println(sum)
}
