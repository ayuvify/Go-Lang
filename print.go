package main

import "fmt"

func main() {
	
	/*
	0) How to Print Something
	    fmt.Println("Hello World")
	*/ 

    /*
	1) To Store String Variable
        var username string = "Ayush"
	    fmt.Println(username)
	    fmt.Printf("Variable is of type: %T \n", username) //To know data type
	*/
	
	
	/*
	2) To Store Boolean 
	    var isLoggedIn bool = true
	    fmt.Println(isLoggedIn)
	    fmt.Printf("Variable is of type: %T \n", isLoggedIn) //To know data type
	*/

	
	/*
	3) To Store Int with different type of memory capacity
	    var smallval uint8 = 225
	    fmt.Println(smallval)
	    fmt.Printf("Variable is of type: %T \n", smallval) //To know data type
	*/

	
	/*
	4) To Declare A Varaiable Without Value
	    var smallDec int
	    fmt.Println(smallDec)f
	    fmt.Printf("Variable is of type: %T \n", smallDec)
	*/


	/* 
	5) Implicit Type Conversion
	    var username = "Ayush"
	    fmt.Println(username)
	    fmt.Printf("Variable is of type: %T \n", username)
	*/

	/*
	6) No Var Declaration But It can be declare only under a function
	    noOfAttempt := 5600
	    fmt.Println(noOfAttempt)
	*/

	/*
	7) Constant and Public Declaration
	const letter = "Private"
	const Letter = "Public" // For Every Public variable or constant need the name start with Capital Letter
	fmt.Println("This is not a public constant", letter);
	fmt.Println("This is a public constant", Letter);
	*/

	var username string = "Ayush"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)
}