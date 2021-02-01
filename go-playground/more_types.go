package main

import "fmt"

func pointers() {
	i, j := 42, 2701
	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j
}

type Vertex struct {
	// X int
	// Y int
	X, Y float64
}

var (
	v1 = Vertex{1, 2}  // has type Vertex
	v2 = Vertex{X: 1}  // Y:0 is implicit
	v3 = Vertex{}      // X:0 and Y:0
	p  = &Vertex{1, 2} // has type *Vertex
)

func iArray() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}

func tSlice() {
	var sNum []int
	sNum = append(sNum, 1, 2, 3)

	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
	for i, v := range pow {
		// for i := range pow {  // Only index
		// for _, value := range pow  // I need value only
		fmt.Printf("2**%d = %d\n", i, v)
	}
}

func tMap() {
	var m = map[string]Vertex{
		"Bell Labs": Vertex{40.68433, -74.39967},
		"Google":    Vertex{37.42202, -122.08408},
	}
	m["Google"] = Vertex{1, 3}
	delete(m, "Google")
	v, ok := m["Google"]  // 0 false
}

func main() {
	pointers()
	v := Vertex{1, 2}
	fmt.Println(v)
	v.X = 4
	fmt.Println(v.X)

	p := &v
	p.X = 1e9
	fmt.Println(v)

	fmt.Println(v1, p, v2, v3)

	iArray()

	// slice
	names := [4]string{"Thomas", "Parker", "Andre", "Henry"}
	a := names[0:2] // [Thomas Parker]
	b := names[1:3] // [Parker Andre]

	b[0] = "Snow"
	fmt.Println(a, b)  // [Thomas Snow] [Snow Andre]
	fmt.Println(names) // [Thomas Snow Andre Henry]

	s := []int{2, 3, 5, 7, 11, 13}
	s = s[1:4] // [3, 5, 7]
	s = s[:2]  // == s[0:2]  // [3, 5]
	s = s[1:]  // == s[1:2]  // [5]

	// length and capacity
	s2 := []int{2, 3, 5, 7, 11, 13} // len=6 cap=6
	s2 = s2[:0]                     // len=0 cap=6
	s2 = s2[:4]                     // len=4 cap=6
	s2 = s2[2:]                     // len=2 cap=4

	var s3 []int // Zero value: len=0 cap=0 []

	// make function
	s4 := make([]int, 5)    // len=5 cap=5 [0 0 0 0 0]
	s5 := make([]int, 0, 5) // len=0 cap=5 []
	s6 := s5[:2]            // len=2 cap=5 [0 0]
	s7 := s6[2:5]           // len=3 cap=3 [0 0 0]

	tSlice()
	tMap()

}
