package main

import "fmt"

func dataTypes() {

	var age uint8 = 28

	//* dec variable short way
	/* var (
		name   string  = "canberk"
		adress string  = "polatlı"
		no     uint8   = 10


		height float32 = 183.5
	) */

	//* Other way for dec

	// var name, adress, no, height = "canberk", "polatlı", 10, 183.5

	//* best short way for dec

	name, adress, no, height := "canberk", "polatlı", 10, 183.5

	fmt.Println(name)
	fmt.Println(adress)
	fmt.Println(no)
	fmt.Println(height)

	fmt.Printf("%T\n", age)
	
	var zeroTYpe int //? string --> "" int --> 0 

	fmt.Println(zeroTYpe) //! zero value
}

func main() {

	var sayHello string = "hello world"

	fmt.Println(sayHello)

	var isLogin bool

	isLogin = true
	fmt.Println(isLogin)

	//* declaration and assign asame time

	name := "canberk koç"
	age := 28

	fmt.Println(name, age)

	//! we can assign different value but we can't change the types

	age = 30
	fmt.Println(name, age)

	//* we can print variables type '%T'

	fmt.Printf("%T\n", name)
	fmt.Printf("%T\n", age)
	fmt.Printf("%T\n", isLogin)

	//? check the golang data types on documentaion

	dataTypes()

}
