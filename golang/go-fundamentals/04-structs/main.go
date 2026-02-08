package main

import "fmt"

// Structs
type car struct {
	brand   string
	model   string
	doors   int
	mileage int
}

// Nested Structs
// The fields of a struct can be accessed using the dot . operator.
type car2 struct {
	brand      string
	model      string
	doors      int
	mileage    int
	frontWheel wheel
	backWheel  wheel
}

type wheel struct {
	radius   int
	material string
}

// Anonymous Structs
// An anonymous struct is just like a normal struct, but it is defined without a name and therefore cannot be referenced elsewhere in the code.
// If a struct is only meant to be used once, then it makes sense to declare it in such a way that developers down the road won’t be tempted to accidentally use it again.
/*
	myCar := struct {
		brand string
		model string
	} {
		brand: "Toyota",
		model: "Camry",
	}
*/
type car3 struct {
	brand   string
	model   string
	doors   int
	mileage int
	// wheel is a field containing an anonymous struct
	wheel struct {
		radius   int
		material string
	}
}

// Embedded Structs
type car4 struct {
	brand string
	model string
}

type truck struct {
	// "car" is embedded, so the definition of a
	// "truck" now also additionally contains all
	// of the fields of the car struct
	car4
	bedSize int
}

// Embedded vs. Nested
// Unlike nested structs, an embedded struct's fields are accessed at the top level like normal fields.
// Like nested structs, you assign the promoted fields with the embedded struct in a composite literal.
func embeddedvsnested() {
	lanesTruck := truck{
		bedSize: 10,
		car4: car4{
			brand: "Toyota",
			model: "Camry",
		},
	}

	fmt.Println(lanesTruck.brand) // Toyota
	fmt.Println(lanesTruck.model) // Camry
}

// Struct Methods
// While Go is not object-oriented, it does support methods that can be defined on structs.
// Methods are just functions that have a receiver. A receiver is a special parameter that syntactically goes before the name of the function.
type rect struct {
	width  int
	height int
}

// area has a receiver of (r rect)
// rect is the struct
// r is the placeholder
func (r rect) area() int {
	return r.width * r.height
}

func structmethods() {
	var r = rect{
		width:  5,
		height: 10,
	}

	fmt.Println(r.area())
}

// Memory Layout
// In Go, structs sit in memory in a contiguous block, with fields placed one after another as defined in the struct.
type contact struct {
	//32 bit - 4 bytes
	sendingLimit int32
	//32 bit - 4 bytes
	age int32
	//pointer size (system in bytes) + int size (system in bytes)
	//32-bit system: pointer size (4 bytes) + int size (4 bytes) = 8 bytes
	//64-bit system: pointer size (8 bytes) + int size (8 bytes) = 16 bytes
	userID string
}

type perms struct {
	permissionLevel int
	canSend         bool
	canReceive      bool
	canManage       bool
}

func main() {
	myCar := car2{}
	myCar.frontWheel.radius = 5
}
