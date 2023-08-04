package main

import("fmt")

func main(){

	// To Store String
	fmt.Println("String Variable")
	var username string = "Ayush"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T", username) //To know data type
	
	// To Store Boolean
	fmt.Println("Boolean Variable")
	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T", isLoggedIn) //To know data type
	
	// To Store Int with different type of memory capacity
	fmt.Println("Int Variable With Different Memory Capacity")
	var smallval uint8 = 225
	fmt.Println(smallval)
	fmt.Printf("Variable is of type: %T", smallval) //To know data type
}