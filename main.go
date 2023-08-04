package main

import (
	"fmt"
)

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


	/*
    11) Array
	
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

	/*
	12) Append Function

	var fruit = []string{"Apple", "Tomato", "Peach"}
	fruit = append(fruit, "Mango", "Banana")
	fmt.Println(fruit)

	// With Starting Point
	fruit = append(fruit[1:])
	fmt.Println(fruit)
	
	// With Ending Point
	fruit = append(fruit[:3])
	fmt.Println(fruit)

	// With Both Starting and Ending Point
	fruit = append(fruit[1:3])
	fmt.Println(fruit)
	*/

	/*
	13) Make Function
	max := make([]int, 4)
	max[0] = 234
	max[1] = 945
	max[2] = 465
	max[3] = 867
	fmt.Println(max)
	*/

	/*
	14) Sort Function
	max := make([]int, 4)
	sort.Ints((max))
	fmt.Println(max)
	*/

	/*
	15) Sort Check Function
	max := make([]int, 4)
	fmt.Println(sort.IntsAreSorted(max))
	*/

	/*
	16) Example of Sort Function
	var course = []string{"react","js","swift","python","ruby"}
	fmt.Println(course)
	index := 2
	course = append(course[:index], course[index+1:]...)
	fmt.Println(course)
	*/

	/* 17) Maps 
	lang := make(map[string]string)

	lang["JS"] = "Javascript"
	lang["RB"] = "Ruby"
	lang["PY"] = "Python"
	fmt.Println("List of all languages: ", lang)
	
	fmt.Println("JS Full Form: ", lang["JS"])
	
	delete(lang, "RB")
	fmt.Println("List of all languages: ", lang)

	// loops on maps

	for key, value := range lang{
		fmt.Printf("For key %v, value is %v\n", key, value)
	}

	for _, value := range lang{
		fmt.Printf("For key v, value is %v\n", value)
	}*/

	/* 18) Structs - Type of Class 
	    //no inheritance in golang; no super or parent
		
		detail := User{"Ayush", "ayush@go.dev", true, 16}
		fmt.Println(detail)
		fmt.Printf("detail of me: %+v\n", detail)
		fmt.Printf("name is: %v and email is %v", detail.Name, detail.Email)*/

		fmt.Print("Casual Print")

	}

	/* 18) Structs - Type of Class
	type User struct {
		Name	string
		Email	string
		Status	bool
		Age	int
	} */
