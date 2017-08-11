package main

import "fmt"

func main() {
	mapTest := make(map[string]int)

	mapTest["kvp"] = 42
	fmt.Println("The value:", mapTest["kvp"])

	mapTest["kvp"] = 48
	fmt.Println("The value:", mapTest["kvp"])

	delete(mapTest, "kvp")
	fmt.Println("The value:", mapTest["kvp"])

	a, ok := mapTest["kvp"]
	fmt.Println("The value:", a, "Present?", ok)
}
