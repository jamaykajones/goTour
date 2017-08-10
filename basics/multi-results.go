package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swap("go?", "how's") // how's go?
	fmt.Println(a, b)
}
