package main

import "fmt"

func main() {
	var a [2]string //array 'a' has 2 values of data
	a[0] = "Hello"  //set first value as 'hello'
	a[1] = "World"  //set second value as 'world'
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13} //array named 'primes'
	fmt.Println(primes)
}
