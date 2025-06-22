package main

import "fmt"

type Rectangle struct {
	length float64
	width  float64
}

// Method with value receiver
func (r Rectangle) Area() float64 {
	return r.length * r.width
}

// Method with pointer receiver
func (r *Rectangle) Scale(factor float64) {
	r.length = r.length * factor
	r.width = r.width * factor
}

func main() {
	r := Rectangle{10, 20}
	area := r.Area()
	fmt.Println(area) // 200

	r.Scale(2)
	area = r.Area()
	fmt.Println(area) // 800

	var myFloat MyFloat
	fmt.Println(myFloat.Something()) // Something from MyFloat
}

type MyFloat float64

func (MyFloat) Something() string {
	return "Something from MyFloat"
}
