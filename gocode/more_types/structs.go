package main

import "fmt"

type Vertex struct { //collection of fields
	X, Y int
}

var (
	v1 = Vertex{8, 9}
	v2 = Vertex{X: 1} //Y:0 is implicit
	v3 = Vertex{}
	p  = &Vertex{1, 2}
)

func main() {
	fmt.Println(v1, p, v2, v3)
}
