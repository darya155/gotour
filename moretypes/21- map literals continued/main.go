package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"Bell Labs": {40.66, -75.785},   //without Vertex name
	"Google":    {37.555, -125.457}, //withoun Vertex name
}

func main() {
	fmt.Println(m)
}
