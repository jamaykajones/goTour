package main

import "fmt"

func main() {
	sum := 1
	for sum < 1000 { //for w/out init and post statement
		sum += sum
	}
	fmt.Println(sum)
}
