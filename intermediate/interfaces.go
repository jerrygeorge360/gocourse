package intermediate

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

func (r rect) area() float64 {
	return r.height * r.width
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (r rect) perimeter() float64 {
	return 2 * (r.height + r.width)
}

func (c circle) perimeter() float64 {
	return math.Pi * c.radius * 2
}

func (c circle) diameter() float64 {
	return 2 * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perimeter())
}

func printType(i interface{}) {
	switch i.(type) {
	case int:
		fmt.Println("Type:int")
	case string:
		fmt.Println("Type:string")
	default:
		fmt.Println("Type:unknown")
	}
}
func myPrinter(i ...interface{}) {
	for _, v := range i {
		fmt.Println(v)
	}
}
func main() {
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}
	measure(r)
	measure(c)
	printType("ultimate")
	myPrinter("Hello", 'r', 54)
}
