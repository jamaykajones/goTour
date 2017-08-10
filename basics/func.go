package main

import "fmt"

func add(x int, y int) int { // take 2 arguments
	return x + y
}

func main() {
	fmt.Println(add(12, 13))
}
