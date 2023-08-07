package main
import("fmt")

func main(){

	// To Store String
	fmt.Println("String Variable")
	var username string = "Ayush"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T", username) // %T - a Go-syntax representation of the type of the value
	
	// To Store Boolean
	fmt.Println("Boolean Variable")
	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T", isLoggedIn)
	
	// To Store Int with different type of memory capacity
	fmt.Println("Int Variable With Different Memory Capacity")
	var smallval uint8 = 225
	fmt.Println(smallval)
	fmt.Printf("Variable is of type: %T", smallval)

	// To Declare A Varaiable Without Value
	var smallDec int
	fmt.Println(smallDec)
	fmt.Printf("Variable is of type: %T \n", smallDec)
}