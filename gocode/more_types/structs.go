package main

import "fmt"

type Vertex struct { //collection of fields
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)
}
