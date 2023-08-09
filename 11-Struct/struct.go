package main

import "fmt"

func main() {

	// Structs - Type of Class
	//no inheritance in golang; no super or parent

	detail := User{"Ayush", "ayush@go.dev", true, 16}
	fmt.Println(detail)
	fmt.Printf("detail of me: %+v\n", detail)
	fmt.Printf("name is: %v and email is %v", detail.Name, detail.Email)
}

// User Define For Struct
type User struct {
	Name	string
	Email	string
	Status	bool
	Age	int
}