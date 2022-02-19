package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {

	fmt.Println(sum(3, 7))
	sum1(3, 87)
	hello("canberk", 28) //! function çalıştırırkne verilen değerler argument

	fmt.Println(result(4))

	var time time.Time = time.Now() //* Now() --> Method

	fmt.Println(time)

	fmt.Print("please enter the note: ")
	reader := bufio.NewReader(os.Stdin)

	value, _ := reader.ReadString('\n') //* _ blank identifier

	fmt.Println(value)

	bölüm, kalan := divide(104, 5)

	fmt.Println(bölüm, kalan)

}

func sum(a int, b int) int { //* a ve b parametre

	return a + b
}

func sum1(a, b int) {

	fmt.Println(a + b)
}

func hello(name string, age int) {

	fmt.Printf("adım %s , yasım %d \n", name, age)

}

func result(grade int) string {

	if grade >= 50 {
		return "pass"
	}

	return "not pass"

}

func divide(bölünen int, bölen int) (bölüm int, kalan int) {

	bölüm = bölünen / bölen

	kalan = bölünen % bölen

	return bölüm, kalan

}
