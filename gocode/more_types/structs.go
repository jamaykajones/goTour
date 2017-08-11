package main

import "fmt"

type Vertex struct { //collection of fields
	X int
	Y int
}

func main() {
	fmt.Println(Vertex{1, 2})
}
