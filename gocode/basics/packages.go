package main

import (
	"fmt"
	"math/rand"
	"time"
)

func rando(min, max int) int {
	return rand.Intn(max - min)
}

func main() {
	rand.Seed(time.Now().Unix())
	myRand := rando
	fmt.Println("My favorite number is", myRand(1, 10))
}
