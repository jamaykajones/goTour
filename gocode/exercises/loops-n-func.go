package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	for i:=0; i < 10; i++ {
		z := float64(1)
	    z = z - (z*z - x)/(2 *z)
	}
	return x
}

func main() {
	fmt.Println(Sqrt(16))
	fmt.Printf("Sqrt of 16 is: %v Am I close?", math.Sqrt(16))
}