package main

import (
	"errors"
	"fmt"
)

func main() {
	// var myString string = "A"
	// fmt.Println(myString)
	//rune are ascii data type
	// var myRune rune = 'A'
	// fmt.Println(myRune)
	//inferred type
	// v1 := 1
	// v2, v3 := 2, 3
	// fmt.Println(v1, v2, v3)
	//constants
	// const NAME string = "Aditya"
	// fmt.Printf("My name is " + NAME)
	//fuctions
	num1 := 20
	num2 := 0
	result, remainder, err := myFuction(num1, num2)
	// if err != nil {
	// 	fmt.Print(err.Error())
	// } else if remainder == 0 {
	// 	fmt.Printf("The result of the integer division is %v", result)
	// } else {
	// 	fmt.Printf("The result of the integer division is %v with remainer %v", result, remainder)
	// }

	//we can use switch case in similar manner
	switch {
	case err != nil:
		fmt.Print(err.Error())
	case remainder == 0:
		fmt.Printf("The result of the integer division is %v", result)
	default:
		fmt.Printf("The result of the integer division is %v with remainer %v", result, remainder)
	}

}

func myFuction(n1 int, n2 int) (int, int, error) {
	var err error
	if n2 == 0 {
		err = errors.New("cannot divide by 0")
		return n1, n2, err
	}
	var result int = n1 / n2
	var remainder int = n1 % n2
	return result, remainder, err
}
