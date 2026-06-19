package main

import "fmt"

// A pointr is a variable that stores the memory address of another variable
func main() {
	b := 255

	var a *int = &b

	fmt.Printf("Type of a is %T\n", a)
	fmt.Println("address of b is", a)
	fmt.Println(*a)

	c := 25
	var d *int
	if d == nil {
		fmt.Println("d is", d)
		d = &c
		fmt.Println("d after initialization is", d)
	}

	// Creating pointers using the new function
	size := new(int)
	fmt.Printf("Size value is %d, type is %T, address is %v\n", *size, size, size)
	*size = 85
	fmt.Println("New size value is", *size)
}
