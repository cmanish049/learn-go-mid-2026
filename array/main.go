package main

import "fmt"

func main() {
	var a [3]int

	fmt.Println(a)

	a[0] = 1
	a[1] = 2
	a[2] = 3

	fmt.Println(a)

	b := [3]int{2, 3, 4}

	fmt.Println(b)

	c := [...]string{"USA", "China", "India", "Germany", "France"}
	d := c // a copy of a is assigned to b
	d[0] = "Singapore"
	fmt.Println("c is ", c)
	fmt.Println("d is ", d)

	for _, v := range c {
		fmt.Println(v)
	}
}
