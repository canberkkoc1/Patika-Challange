package main

import "fmt"

//! when we declaration package variable (global) this variable uses the memory store until the program finish. But when we declaration in func variable use the memory store only func run

var pack string = "global scope"

//! you can not short declaration outside function...
//* -beacuse we do not said what it is this variables. We will say in func this is variables...
/*

packEr := 23

*/

func typeConv() {

	age := 10
	w := 23.8
	//* "%v" is print the variables value only run Printf

	fmt.Printf("%v %T\n", age, age)

	//! you can not this beacuse 2 variables different type
	// fmt.Println(age + w)

	//* type conversion

	fmt.Println(int(w) + age)

	//! we do not change the type of the variable just convert when you look below log w type float64

	fmt.Printf("%v %T\n", w, w)

	//! string type cannot convert int
	/*
		x := 10
		y := "10"

		fmt.Println(x + int(y))
	*/

	num := 106
	//* string(106) is charcode
	str := string(106)

	fmt.Printf("%v %T\n", num, num)
	fmt.Printf("%v %T\n", str, str)

}

func scopes() {

	name := "in scopes"

	fmt.Println(name, pack)

	//	fmt.Println(packEr)

	if true {
		var block string = "block scope"
		fmt.Println("in if block", block)
	}

	//! error = undeclarate name (block)
	// fmt.Println(block)
}

func main() {

	scopes()

	//* declaration

	var name string

	//* Assign

	var age int = 28

	//* Initialization

	var check bool = true

	fmt.Println(age)
	fmt.Println(name)
	fmt.Println(check)

	fmt.Println("*****************************")

	typeConv()
}
