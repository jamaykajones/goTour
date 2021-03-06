package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("can't Sqrt a negative number: %v", float64(e)) 
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	} else {
		z :=  float64(1)
		for i:=0; i < 1000; i++ {
			z = z - (z*z - x)/(2 *z)
		}	
		return z, nil
	}
}

func main() {
	fmt.Println(Sqrt(16))
	fmt.Println(Sqrt(-16))
}
