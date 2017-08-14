package main

import "fmt"

type Dest struct {
	Lat, Long float64
}

var m = map[string]Dest{"Bell Labs": Dest{40.68433, -74.39967},
	"Google": Dest{37.42202, -122.08408},
}

func main() {
	fmt.Println(m)
}
