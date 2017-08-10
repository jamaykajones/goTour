package main

import "fmt"

func add(x, y int) int { // if arguments are the same type than all but the last can be omitted
	return x + y
}

func main() {
	fmt.Println(add(23, 32)) //55
}
