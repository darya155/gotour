package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	p := &v
	p.X = 1e9
	fmt.Println("1 -", v)
	fmt.Println("2 -", p.Y)
	fmt.Println("3 -", v.Y)
	fmt.Println("4 -", p)
	fmt.Println("5 -", *p)
	fmt.Println("6 -", &v)
	fmt.Println("7 -", &p)
}
