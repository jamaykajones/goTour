package main

import (
	"fmt"
	"math/rand"
	"time"
)

func rando(min, max int) int {
	return rand.Intn(max-min) + min
}

func main() {
	rand.Seed(time.Now().Unix()) //time is always '2009-11-10 23:00:00 UTC' in Go
	myRand := rando
	fmt.Println("My favorite number is", myRand(1, 10))
}
