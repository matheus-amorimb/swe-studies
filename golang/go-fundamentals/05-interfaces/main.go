package main

import "fmt"

// Interfaces
// Interfaces are implemented implicitly.
// A type never declares that it implements a given interface.
// Implicit interfaces decouple the definition of an interface from its implementation.
type shape interface {
	area() float64
	perimeter() float64
}

type rect struct {
	width, height float64
}

func (r rect) area() float64 {
	return r.width * r.height
}
func (r rect) perimeter() float64 {
	return 2*r.width + 2*r.height
}

func printShapeData(s shape) {
	fmt.Printf("Area: %v - Perimeter: %v\n", s.area(), s.perimeter())
}

// Multiple Interfaces
// A type can implement any number of interfaces in Go.

// Name Your Interface Parameters
type CopierNotNamed interface {
	Copy(string, string) int
}

// named parameters
type CopierNamed interface {
	Copy(sourceFile string, destinationFile string) (bytesCopied int)
}

// Type Assertions
type shape2 interface {
	area() float64
}

type circle2 struct {
	radius float64
}

func (c circle2) area() float64 {
	return 3.14 * c.radius
}

func typeassertions(s shape2) {
	c, ok := s.(circle2)
	if !ok {
		// log an error if s isn't a circle
		fmt.Println("s is not a circle")
	}
	radius := c.radius
	fmt.Println(radius)
}

func main() {
	r := rect{
		width:  4.0,
		height: 6.0,
	}
	printShapeData(r)
}

// Type Switches
// A type switch makes it easy to do several type assertions in a series.
func printNumericValue(num interface{}) {
	switch v := num.(type) {
	case int:
		fmt.Printf("%T\n", v)
	case string:
		fmt.Printf("%T\n", v)
	default:
		fmt.Printf("%T\n", v)
	}
}

type Formatter interface {
	Format() (formattedString string)
}

type PlainText struct {
	message string
}

func (p PlainText) Format() (formattedString string) {
	return p.message
}

type c struct {
	message string
}

func (b Bold) Format() (formattedString string) {
	return fmt.Sprint("*%s*", b.message)
}

type Code struct {
	message string
}

func (c Code) Format() (formattedString string) {
	return fmt.Sprint("`%s`", c.message)
}

// Don't Touch below this line

func SendMessage(formatter Formatter) string {
	return formatter.Format() // Adjusted to call Format without an argument
}
