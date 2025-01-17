package main

import "fmt"

func main() {

	// Generic For Loop
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// For Loop works as while loop
	for i := 5; i < 10; {
		fmt.Println(i)
		i++
	}

}
