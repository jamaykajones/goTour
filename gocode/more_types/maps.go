package main

import "fmt"

type Dest struct {
	Lat, Long float64
}

var m map[string]Dest

func main() {
	m = make(map[string]Dest)
	m["Bell Labs"] = Dest{
		40.68433, -74.39967,
	}
	fmt.Println(m)
}
