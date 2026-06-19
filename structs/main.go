package main

import "fmt"

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

	emp2 := Employee{"Subash", "Chaudhary", 32, 140, Address{"Dhangadhi", "Sudurpaschim"}}

	fmt.Println("Employee 1", emp1)
	fmt.Println("Employee 2", emp2)

	fmt.Println("First name of employee 1", emp1.firstname)
	fmt.Println("Last name of employee 1", emp1.lastname)
	fmt.Println("Age of employee 1", emp1.age)
	fmt.Println("Salary of employee 1", emp1.salary)

	fmt.Printf("Address of employee 1 is %s, %s\n", emp1.city, emp1.state)
}
