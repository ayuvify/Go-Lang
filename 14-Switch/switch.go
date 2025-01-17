package main

import "fmt"

func main() {
	var dx int
	fmt.Println("Input A Number: ")
	fmt.Scan(&dx)
	switch dx {
	case 1:
		fmt.Print("One")
	case 2:
		fmt.Print("Two")
	case 3:
		fmt.Print("Three")
	default:
		fmt.Print("More than 3")
	}
}
