package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {

	// sum, div, mult, sub := allOperations(3, 2)

	// fmt.Println(sum, div, mult, sub)

	// findNumber()

	/* fmt.Print("please enter the note: ")

	result, _ := grade()

	if result >= 50 {
		fmt.Println("geçtin")
	} else {

		fmt.Println("kaldın")
	}
	*/

	target := findNum(1, 100)

	fmt.Println("1 ile 100 arasında sayı bulmaya çalış")

	reader := bufio.NewReader(os.Stdin)

	for attemps := 0; attemps < 10; attemps++ {
		fmt.Println(10-attemps, "hakkınız kaldı")
		fmt.Print("guess : ")

		input, err := reader.ReadString('\n')

		input = strings.TrimSpace(input)

		num, err := strconv.Atoi(input)

		if err != nil {
			fmt.Println(err)
		}

		if num == target {
			fmt.Println("sayıyı buldunuz.")
			break
		} else if num <= target {

			fmt.Println("daha büyük bir sayı giriniz.")
		} else {
			fmt.Println("daha küçük bir sayı giriniz.")

		}

	}

}

func allOperations(a, b int) (int, int, int, int) {

	sum := a + b
	divide := a / b
	mult := a * b
	subtraction := a - b

	return sum, divide, mult, subtraction

}

/* func grade() string {

	var num int
	fmt.Print("please enter the note: ")

	fmt.Scanf("%d", num)

	if num > 50 {
		return "geçtiniz"
	}
	return "kaldınız"

} */

func grade() (int, error) {

	reader := bufio.NewReader(os.Stdin)

	input, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println(err)
	}

	//* boşluk silme
	input = strings.TrimSpace(input)

	num, err := strconv.Atoi(input) //* string value to integer

	if err != nil {
		fmt.Println(err)
	}

	return num, nil

}

/* func findNumber() {

	var num int
	value := rand.Intn(100-1) + 1
	fmt.Println(value)
	for i := 0; i < 10; i++ {

		fmt.Print("please guess number : ")
		fmt.Scanf("%d", num)

		if num == value {
			fmt.Println("yes you guess the number")
			break
		}

	}

}
*/

func findNum(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}
