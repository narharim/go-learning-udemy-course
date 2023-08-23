package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

type shape interface {
	Area() float64
}

func calculateArea(s shape) float64 {
	return s.Area()
}

type circle struct {
	Radius float64
}

type square struct {
	length float64
}

func (c circle) Area() float64 {
	return 2 * 3.14 * c.Radius * c.Radius
}

func (r square) Area() float64 {
	return 4 * r.length
}

func main() {
	var b bytes.Buffer

	b.Write([]byte("Hello"))

	fmt.Fprintf(&b, " World\n")
	io.Copy(os.Stdout, &b)

	circle := circle{Radius: 15}
	square := square{length: 4}

	fmt.Printf("Area of the shape: %.2f\n", calculateArea(square))
	fmt.Printf("Area of the shape: %.2f\n", calculateArea(circle))

}
