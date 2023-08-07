package main

import "fmt"

func main() {

	// To Declare A Varaiable Without Value
	var smallDec int
	fmt.Println(smallDec)
	fmt.Printf("Variable is of type: %T \n", smallDec)

	// Implicit Type Conversifmt
	var username = "Ayush"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	// No Var Declaration But It can be declare only under a function
	noOfAttempt := 5600
	fmt.Println(noOfAttempt)

	// Constant and Public Declaration
	const letter = "Private"
	const Letter = "Public" // For Every Public variable or constant need the name start with Capital Letter
	fmt.Println("This is not a public constant", letter);
	fmt.Println("This is a public constant", Letter);
}