package main

import (
	"fmt"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) add() float64 {
	return v.X + v.Y
}

func main() {
	fmt.Println("Hello Go!")
	v := Vertex{10.1, 11.9}
	fmt.Println(v.add())
}

func add(i int, j int) int {
	return i + j
}
