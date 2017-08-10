package main

import "fmt"

var i, j int = 1, 2 //one initializer per variable, type omitted

func main() {
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)
}
