package main

import "fmt"

var y = "canb"

const num1 = 1

func main() {

	x := 50

	//* assigment statement
	//x = x + 10

	x += 10 //*assigmetn operation (atama ve işlem aynı anda)

	fmt.Printf("%T %v\n", x, x)

	F := -40

	K := float64(F-32)/1.8 + 273
	fmt.Printf("%T %v\n", K, K)

	//! constant variables shadowing

	var y = "berk"
	fmt.Printf("%T %v\n", y, y)

	const num1 = 3223
	fmt.Printf("%T %v\n", num1, num1)

	t, g := "a", "b"

	fmt.Printf("%T %v\n", t == g, t == g)
	fmt.Printf("%T %v\n", t != g, t != g)
	fmt.Printf("%T %v\n", t < g, t < g)
	fmt.Printf("%T %v\n", t > g, t > g)

}
