package main

import (
	"fmt"
	"math"
	"reflect"
)

// Nil interface
type Template interface {}

type Shaper interface {
    Area() float64
}
// Tao ra struct cycle, rectangle and va non-struct
type Circle struct {
	radius float64
}
type Rectangle struct {
	x int
	y int
}
type myfloat float64
// Implemetn interface
func (c *Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}
func (r *Rectangle) Area() int {
	return r.x * r.y 
}
func (f myfloat) Area() string {
	return "Non-Struct type can implemetn Interface"
}
// Su dung nil interface de nhan vao unknown argument
func PrintTemplate(t Template) string {
	switch reflect.TypeOf(t) {
    case reflect.TypeOf(&Circle{}):
        return "Circle"
    case reflect.TypeOf(&Rectangle{}):
        return "Rectangle"
    case reflect.TypeOf(myfloat(0)):
        return "Float"
    default:
        return "Other Type"
    }
}

func main() {
	// Testcase cho Circle
	circle := &Circle{radius: 5}
	resultCircle := circle.Area()
	fmt.Printf("Area of Circle: %.2f\n", resultCircle)

	// Testcase cho Rectangle
	rectangle := &Rectangle{x: 4, y: 6}
	resultRectangle := rectangle.Area()
	fmt.Printf("Area of Rectangle: %d\n", resultRectangle)

	// Testcase cho myfloat
	var nonStructType myfloat = 10.0
	resultNonstruct := nonStructType.Area()
	fmt.Printf("Area of myfloat: %s\n", resultNonstruct)

	// Test nil interface
	cir := PrintTemplate(circle)
	fmt.Println(cir)
	rec := PrintTemplate(rectangle)
	fmt.Println(rec)
	fl := PrintTemplate(resultNonstruct)
	fmt.Println(fl)
    // Test case cho non-struct type
    fmt.Println(nonStructType.Area()) 
	fmt.Println("Nil Interface is used in order to receiver unknown argument in function")
	fmt.Println("Study Inteface")
}

