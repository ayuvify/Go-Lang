package main

import "fmt"

func main()  {
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
}