package main

import "fmt"

func calculateBill(price int, quantity int) int {
	var totalPrice = price * quantity

	return totalPrice
}

func calculateBill2(price, quantity int) int {
	return price * quantity
}

// Multiple return
func rectangleProps(length, width float64) (float64, float64) {
	var area = length * width

	var perimeter = (length * width) * 2

	return area, perimeter
}

func rectangleProps2(length, width float64) (area, perimeter float64) {
	area = length * width

	perimeter = (length * width) * 2

	return
}

// Blank identifier

func rectanglePropsBlankIdentifier(length, width float64) (float64, float64) {
	var area = length * width

	var perimeter = (length * width) * 2

	return area, perimeter
}

func main() {
	area, _ := rectanglePropsBlankIdentifier(10.8, 5.2)

	fmt.Printf("Area %f ", area)
}
