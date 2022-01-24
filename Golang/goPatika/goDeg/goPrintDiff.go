package main

import "fmt"

//! GLoabal variables should be start uppercase

var VisNum = 128

func main() {

	/*
		//?	check the this web page https://pkg.go.dev/fmt
		//! and watch this video for golang best preactice https://www.youtube.com/watch?v=MzTcsI6tn-0

	*/

	//* Printf is format type example below

	name := "ck"

	fmt.Printf("my name is %v\n", name)

	//* %b print the 2 based
	num := 20

	fmt.Printf("%b", num)

	//! we can set whitespace print but println we dont need to set whitespace.

	fmt.Print("my name is ", name, "\n")
	fmt.Println("my name is", name)

	//! visibility

	fmt.Println(VisNum)

}
