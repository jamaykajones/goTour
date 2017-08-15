package main

import "fmt"

type I interface {
	M()
}

func main() {
	var i Ipackage main
	
	import "fmt"
	
	func main() {
		var i interface{} // empty interface
		describe(i)
	
		i = 42
		describe(i)
	
		i = "hello"
		describe(i)
	}
	
	func describe(i interface{}) {
		fmt.Printf("(%v, %T)\n", i, i)
	}
	
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
