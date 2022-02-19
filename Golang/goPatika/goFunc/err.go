package main

import (
	"errors"
	"fmt"
	"math"
)

//! go errorleri bir value olarak döner.

func main() {

	result, err := evenNum(9)

	if err != nil {
		fmt.Println(err)
	} else {

		fmt.Println(result)
	}

	root, err := sRoot(-2)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(root)
	}

}

func evenNum(num int) (string, error) { //* bir hata olursa error döndür.

	if num%2 != 0 {
		return "", errors.New("çift sayı girmediniz")
	}

	return "not even", nil //! nil ifadesi başlangıç değeri olmayan ifadelere verilen değer

}

func sRoot(num float64) (float64, error) { //? error bir type

	if num < 0 {
		return 0, errors.New("negatif sayıların karekökü alınamaz.")
	}

	return math.Sqrt(num), nil

}
