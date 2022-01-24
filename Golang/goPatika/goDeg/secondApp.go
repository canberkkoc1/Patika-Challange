package main

import (
	"fmt"
	"strconv"
)

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

}
