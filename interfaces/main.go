package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perimeter() float64
}

type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}

func main() {
	r := rect{width: 10, height: 3}
	c := circle{radius: 5}
	measure(r)
	measure(c)

	myPrinter(1, "aaa", "bbb", "ccc", 1.5, true)

	printType(1)   //Type: Int
	printType("a") // Type: String
	printType(1.5) // Type: Unknown

}

// Another use of interfaces - 1
func myPrinter(m ...interface{}) {
	for _, msg := range m {
		fmt.Println(msg)
	}
}

// Another use of interfaces - 2
func printType(i interface{}) {
	switch i.(type) {
	case int:
		fmt.Println("Type: Int")
	case string:
		fmt.Println("Type: String")
	default:
		fmt.Println("Type: Unknown")

	}
}

func measure(g geometry) {
	fmt.Printf("type of shape: %T\n", g)
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perimeter())

}

func (r rect) area() float64 {
	return r.height * r.width
}

func (r rect) perimeter() float64 {
	return 2 * (r.height + r.width)
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

// func (c circle) diameter() float64 {
// 	return 2 * c.radius
// }
