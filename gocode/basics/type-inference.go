package main

import "fmt"

func main() {
	v := 42 // variable's type is inferred from the value on the right
	x := 42.
	y := "42"
	z := 0.867 + 0.5i
	fmt.Printf("v is of type %T\n", v) // %T displays type
	fmt.Printf("x is of type %T\n", x)
	fmt.Printf("y is of type %T\n", y)
	fmt.Printf("z is of type %T\n", z)
}
