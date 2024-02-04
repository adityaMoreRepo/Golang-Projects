package main

import "fmt"

func main() {
	//tower of honoi problem
	toh(3, 'A', 'B', 'C')
}

func toh(n int, A rune, B rune, C rune) {
	if n == 1 {
		fmt.Printf("Move disc 1 from %c to %c\n", A, C)
		return
	}
	toh(n-1, A, C, B)
	fmt.Printf("Move disc %v from %c to %c\n", n, A, C)
	toh(n-1, B, A, C)
}
