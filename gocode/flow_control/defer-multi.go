package main

import "fmt"

func main() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}

//Deferred function calls are pushed onto a stack.
//Executed in last-in-first-out order
//9, 8, 7, 6, 5, 4, 3, 2, 1, 0
