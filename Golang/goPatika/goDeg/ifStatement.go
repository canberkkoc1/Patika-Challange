package main

import "fmt"

func main() {

	x := 33

	y := -1

	if x%2 == 0 {

		fmt.Println("x çift sayıdır")

	} else {
		fmt.Println("x çift sayı degıldır")

	}

	if y < 0 {
		fmt.Println("sayı negatiftir")
	} else if y%2 == 0 {
		fmt.Println("sayi çift sayıdır")
	} else {
		fmt.Println("sayı tek sayıdır")
	}

}
