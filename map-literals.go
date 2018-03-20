package main

import "fmt"

type Vertex struct {
	Lat, long float64
}

var m = map[string]Vertex{
	"Bell Lab": Vertex{
		40.68433, -74.39967,
	},
	"Google": Vertex{
		37.42202, -122.08408,
	},
}

func main() {
	fmt.Println(m)
}