package main

import "fmt"

func main() {

	// const x int64 = 4

	// fmt.Printf("%T %v\n", x, x)

	//* typless constant

	//? go declarate default type when we declarate typless constant (int)

	const x = 4
	const y = 4.4

	fmt.Printf("%T %v\n", x, x)
	fmt.Printf("%T %v\n", y, y)

	//! normally we can get an error int + float but here const x typeless this is why we don't get an error (const x type conversion)

	fmt.Printf("%T %v\n", x+y, x+y)

	//! go doesnt have prefix (++x) just use postfix (x++)

}
