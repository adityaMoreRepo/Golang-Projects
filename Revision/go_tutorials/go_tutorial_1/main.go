package main

import (
	"fmt"
	"time"
)

func main() {
	//Variables : Initialization, declaration, Assignments and more

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
	// num1 := 20
	// num2 := 0
	// result, remainder, err := myFuction(num1, num2)
	// if err != nil {
	// 	fmt.Print(err.Error())
	// } else if remainder == 0 {
	// 	fmt.Printf("The result of the integer division is %v", result)
	// } else {
	// 	fmt.Printf("The result of the integer division is %v with remainer %v", result, remainder)
	// }

	//we can use switch case in similar manner
	// switch {
	// case err != nil:
	// 	fmt.Print(err.Error())
	// case remainder == 0:
	// 	fmt.Printf("The result of the integer division is %v", result)
	// default:
	// }
	//we also have conditional switch like java
	// switch remainder {
	// case 0:
	// 	fmt.Print("Remainer is zero")
	// case 1, 2, 3:
	// 	fmt.Print("Other valuse")
	// }

	//Arrays
	// var list [3]int32 = [3]int32{1, 2, 3} // initializes array of size lenght 12 Byte = 3 * 4(32 bits) Byte
	// fmt.Println(&list[0])
	// fmt.Println(&list[1])
	//simple shorthand
	// intArr := [3]int32{1, 2, 5}
	// fmt.Print(intArr)
	//we can infer the size by using ...
	// intArr := [...]int32{1, 2, 5}
	// fmt.Print(intArr)
	// slices are just wrappers around the arrays, i.e. with some additional functionality
	//by emiting the size we get a slice. Slices are basically dynamic size arrays.
	// intSlice := []int{1, 2, 10}
	// fmt.Printf("size : %v capacity : %v\n", len(intSlice), cap(intSlice))
	// intNewSlice := append(intSlice, 5)
	// fmt.Printf("size : %v capacity : %v\n", len(intNewSlice), cap(intNewSlice))
	// fmt.Println(intSlice)
	// fmt.Println(intNewSlice)
	//Upending multiple values using spread operator
	// intSlice := []int{1, 2, 10}
	// fmt.Printf("size : %v capacity : %v\n", len(intSlice), cap(intSlice))
	// sliceApp := []int{12, 13}
	// intSlice = append(intSlice, sliceApp...)
	// fmt.Printf("size : %v capacity : %v\n", len(intSlice), cap(intSlice))

	//we can also use make to create slices
	// intSlice := make([]int32, 3, 12)
	// fmt.Printf("size : %v capacity : %v\n", len(intSlice), cap(intSlice))

	//Map
	// var myMap = map[string]int32{"Aditya": 22, "Vishal": 33, "Sheetal": 28}
	//OR
	// myMap2 := make(map[string]int32)
	// fmt.Println(myMap2)
	// fmt.Println(myMap)
	//maps in go can return defualt value of the Value here int32 if the kye is not present, but also return a boolean value
	// fmt.Println(myMap["Aditya"])
	// num, ok := myMap["Pranjal"]
	// if ok {
	// 	fmt.Println(num)
	// } else {
	// 	fmt.Println("Invalid key provided")
	// }
	//delete by key : as it delete by referece of key
	// delete(myMap, "Vishal")
	//iterate through map
	// for name, age := range myMap {
	// 	fmt.Printf("name: %v age: %v\n", name, age)
	// }

	//Iterate through slice
	// stringSlice := []string{"Vikas", "Prerarna", "Amol"}
	// stringSlice = append(stringSlice, "Aditya")
	// for i, name := range stringSlice {
	// 	fmt.Printf("Index: %v name: %v\n", i, name)
	// }
	// for i := 0; i < len(stringSlice); i++ {
	// 	fmt.Printf("Index: %v name: %v\n", i, stringSlice[i])
	// }

	//It is a good practive to preinitiallize the slice for better performance
	n := 1000000
	testSlice := []int{}
	testSlice2 := make([]int, 0, n) //init a slice with size 0 and cap 1000000
	//let's test the time taken to append each slice with million entries
	fmt.Printf("Total time taken to append without preallocation: %v\n", timeLoop(testSlice, n))
	fmt.Printf("Total time taken to append with preallocation: %v\n", timeLoop(testSlice2, n))
}

func timeLoop(slice []int, n int) time.Duration {
	t0 := time.Now()
	for len(slice) < n {
		slice = append(slice, 1)
	}
	return time.Since(t0)
}

// func myFuction(n1 int, n2 int) (int, int, error) {
// 	var err error
// 	if n2 == 0 {
// 		err = errors.New("cannot divide by 0")
// 		return n1, n2, err
// 	}
// 	var result int = n1 / n2
// 	var remainder int = n1 % n2
// 	return result, remainder, err
// }
