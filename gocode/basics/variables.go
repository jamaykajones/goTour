package main

import "fmt"

func main() {
	i, j := 1, 2 //short declaration must be inside func (:=)
	k := 3
	c, python, java := true, false, "no!"

	fmt.Println(i, j, k, c, python, java)
}