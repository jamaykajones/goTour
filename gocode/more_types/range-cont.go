package main

import "fmt"

func main() {
	pow := make([]int, 10)
	for i := range pow {
		pow[i] = 1 << uint(i) // == 2**i
	}
	for _, value := range pow { // drop the index or value intirely
		fmt.Printf("%d\n", value)
	}
}
