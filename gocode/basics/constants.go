package main

import "fmt"

const Pi = 3.14 //cannot use (:=)

func main() {
	World := "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	Truth := true
	fmt.Println("Go rules?", Truth)
}
