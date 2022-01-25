package main

import (
	"fmt"
	"math"
	"strconv"
)

func constant() {

	//! constants are create ccompile time / const --> compile time
	//? variables are create run time / var --> run time

	r := 3.5

	const pi float64 = 3.14

	areaOfCircle := pi * (math.Pow(r, 2))
	// areaOfCircle := 3.14 * (math.Pow(r, 2))

	fmt.Println(areaOfCircle)

	//* this code run because function run in run time

	// var x = math.Pow(3, 4)

	// fmt.Printf("%T %v\n", x, x)

	//! this code give an error because const run from compile time but function (Pow) run from run time

	// const x = math.Pow(3, 4)
	// fmt.Printf("%T %v\n", x, x)

	//! this code give an error same reason above (y is a variable and run from run time )
	// y := 4

	// const x = y

}

func main() {

	x := 75

	var y float64

	y = float64(x)

	fmt.Println(y)
	fmt.Println("****************")

	num1 := 39
	num2 := 57

	fmt.Println(num1, num2)
	fmt.Println("****************")

	num1, num2 = num2, num1

	fmt.Println(num1, num2)

	fmt.Println("****************")

	//! shadowing

	num3 := 54

	if true {

		//* this num3 variable just assign and  num3  = 15 whole func

		num3 = 15

		//* this num3 variable declarate again and just run this block
		// num3 := 10

		num3--

		fmt.Println(num3)
	}

	fmt.Println(num3)

	fmt.Println("****************")

	num4 := 46

	//* here number 46 encode 46 equl = .
	str := string(num4)

	fmt.Println(str)
	fmt.Println("****************")

	//* but here we get "46" in string

	str1 := strconv.Itoa(num4)

	fmt.Println(str1)

	constant()

}
