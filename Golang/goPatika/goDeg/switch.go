package main

import "fmt"

func main() {

	x := 3

	switch x {
	case 6:
		fmt.Println("sayi 6")
	case 5:
		fmt.Println("sayı 5")

	default:
		fmt.Println("sayı ne 6 ne de 5")
	}

}
