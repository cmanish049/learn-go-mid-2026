package main

import (
	"fmt"
	"math"
)

func main() {
	var age int

	fmt.Println("My intial age", age)

	age = 34

	fmt.Println("My age after first assignment", age)

	age = 50

	fmt.Println("My age after second assignment", age)

	var name string = "Manish"

	fmt.Println("Name is ", name)

	// Multiple variable declaration
	var price, quantity int = 5000, 100

	fmt.Println("price is", price, "quantity is", quantity)

	// Shorthand
	count := 10

	fmt.Println("Count = ", count)

	a, b := 10, 20

	fmt.Println(a, b)

	b, c := 30, 40

	fmt.Println(a, b, c)

	e, f := 20.34, 21.32
	d := math.Min(e, f)

	fmt.Println("Minimum : ", d)

	marks := 30

	// marks := "English Marks" as go is strongly typed, variable assigned to a type cannot be assigned as different type

	fmt.Println("Marks : ", marks)
}
