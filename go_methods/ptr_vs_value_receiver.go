//This example shows the difference between value and pointer receiver

package main

import (
	"fmt"
)

type Employee struct {
	name string
	age  int
}

/*
Method with value receiver
*/
func (e Employee) changeName(newName string) {
	e.name = newName
}

/*
Method with pointer receiver
*/
func (e *Employee) changeAge(newAge int) {
	e.age = newAge
}

func main() {
	e := Employee{
		name: "ABC",
		age:  30,
	}
	fmt.Printf("Employee name before change: %s", e.name)
	e.changeName("XYZ")
	fmt.Printf("\nEmployee name after change: %s", e.name)

	fmt.Printf("\n\nEmployee age before change: %d", e.age)
	e.changeAge(35)
	fmt.Printf("\nEmployee age after change: %d\n", e.age)
}
