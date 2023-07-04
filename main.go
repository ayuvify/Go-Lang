package main

import "fmt"

func main() {
	
	/*
	1) How to Print Something 
	    fmt.Println("Hello World")
	*/

    /*
	2) To Store String Variable
        var username string = "Ayush"
	    fmt.Println(username)
	    fmt.Printf("Variable is of type: %T \n", username) //To know data type
	*/
	
	
	/*
	3) To Store Boolean 
	    var isLoggedIn bool = true
	    fmt.Println(isLoggedIn)
	    fmt.Printf("Variable is of type: %T \n", isLoggedIn) //To know data type
	*/

	
	/*
	4) To Store Int with different type of memory capacity
	    var smallval uint8 = 225
	    fmt.Println(smallval)
	    fmt.Printf("Variable is of type: %T \n", smallval) //To know data type
	*/

	
	/*
	5) To Declare A Varaiable Without Value
	    var smallDec int
	    fmt.Println(smallDec)f
	    fmt.Printf("Variable is of type: %T \n", smallDec)
	*/


	/* 
	6) Implicit Type Conversion
	    var username = "Ayush"
	    fmt.Println(username)
	    fmt.Printf("Variable is of type: %T \n", username)
	*/

	/*
	7) No Var Declaration But It can be declare only under a function
	    noOfAttempt := 5600
	    fmt.Println(noOfAttempt)
	*/

	/*
	8) Constant and Public Declaration
	const letter = "Private"
	const Letter = "Public" // For Every Public variable or constant need the name start with Capital Letter
	fmt.Println("This is not a public constant", letter);
	fmt.Println("This is a public constant", Letter);
	*/

	/*
	9) Time Study of GoLang
	present := time.Now()
	create := time.Date(2020, time.April,10,23,23,0,0,time.UTC)

	fmt.Println("Present Date Without Format :",present)
	fmt.Println("Present Date With Format :", present.Format("02-01-2006 15:04:05 Monday"))
	fmt.Println("Created Date :",create)
	*/


	/* 
	10) Pointers

	A) Normal Init A Variable
	    var one int = 2
	    fmt.Println(one)
	
	B) Pointer Init A Variable
	    var ptr *int
	    fmt.Print(ptr)
	
	C) Memory Address
	    num := 12
	    var ptr = &num
	    fmt.Println(ptr)
	*/


	/* 11) Array
	
	A) One Type of Array Init
	    var fl [4]string 
	    fl[0] = "Apple"
	    fl[1] = "Banana"
	    fl[3] = "Cucumber"
	
	    fmt.Println(fl)
	    fmt.Println(len(fl))

	B) Second Type of Array Init
	    var vl = [3]string{"potato","beans","mushroom"}
	    fmt.Println(vl)
	*/

	fmt.Print("Casual Print")
}