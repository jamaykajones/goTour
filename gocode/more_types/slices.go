package main

import "fmt"

func main() {
	primesArray := [6]int{2, 3, 5, 7, 11, 13}

	var sliceOfArray []int = primesArray[1:6]
	fmt.Println(sliceOfArray)
}
