package main

import (
	"fmt"

	"manishdev.com/structs/geometry"
)

type Address struct {
	city  string
	state string
}

type Employee struct {
	firstname string
	lastname  string
	age       int
	salary    int
	Address
}

func (a Address) fullAddress() {
	fmt.Printf("Full Address: %s, %s\n", a.city, a.state)
}

// Methods
func (e Employee) displaySalary() {
	fmt.Printf("Salary of %s is %d\n", e.firstname, e.salary)
}

func (e Employee) changeName(newName string) {
	e.firstname = newName
}

func (e *Employee) changeAge(newAge int) {
	e.age = newAge
}

func main() {
	emp1 := Employee{
		firstname: "Manish",
		lastname:  "Chaudahry",
		age:       34,
		salary:    100,
		Address: Address{
			city:  "Dhangadhi",
			state: "SudurPaschim",
		},
	}

	emp1.displaySalary()
	emp1.changeName("manish")
	emp1.changeAge(35)
	fmt.Println("New Name: ", emp1.firstname)
	fmt.Println("New Age: ", emp1.age)

	emp2 := Employee{"Subash", "Chaudhary", 32, 140, Address{"Dhangadhi", "Sudurpaschim"}}
	emp2.displaySalary()
	emp2.fullAddress()

	fmt.Println("Employee 1", emp1)
	fmt.Println("Employee 2", emp2)

	fmt.Println("First name of employee 1", emp1.firstname)
	fmt.Println("Last name of employee 1", emp1.lastname)
	fmt.Println("Age of employee 1", emp1.age)
	fmt.Println("Salary of employee 1", emp1.salary)

	fmt.Printf("Address of employee 1 is %s, %s\n", emp1.city, emp1.state)

	r := geometry.Rectangle{
		Length: 10,
		Width:  5,
	}

	c := geometry.Circle{
		Radius: 7,
	}

	fmt.Println("Rectangle area:", r.Area())
	fmt.Println("Circle area:", c.Area())
}
