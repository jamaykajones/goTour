package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)
	s = s[:0]
	printSlice(s)

	s = s[:4]
	printSlice(s)

	s = s[2:]
	printSlice(s)
}

func printSlice(s []int) { //PrintSlice prints a slice in a more readable way

	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
