package main

import (
    "fmt"
    "math"
	"reflect"
)

// Interface shape
type Shape interface {
    area() float64
}

// Circle structure
type Circle struct {
    radius float64
}

// Circle area
func (c Circle) area() float64 {
    return math.Pi * c.radius * c.radius
}

type Rectangle struct {
    length float64
    width float64
}

func (r Rectangle) area() float64 {
    return r.length * r.width
}

type Triangle struct {
    length float64
    height float64
}

func (t Triangle) area() float64 {
    return 0.5 * t.length * t.height
}

func main() {
    circle := Circle{radius: 5}
    rectangle := Rectangle{length: 4, width: 6}
    triangle := Triangle{length: 3, height: 8}

    shapes := []Shape{circle, rectangle, triangle}

    for _, shape := range shapes {
        fmt.Printf("Area of %s is : %f\n", reflect.TypeOf(shape), shape.area())
    }
}