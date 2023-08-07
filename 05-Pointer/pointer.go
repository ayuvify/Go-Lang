package main

import "fmt"

func main() {
	// A) Normal Init A Variable
	var one int = 2
	fmt.Println(one)
	
	// B) Pointer Init A Variable
	var ptr *int
	fmt.Println(ptr)
	
	// C) Memory Address
	num := 12
	var ptr1 = &num
	fmt.Println(ptr1)
}