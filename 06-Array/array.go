package main

import "fmt"

func main() {
	// A) One Type of Array Init
	var fl [4]string
	fl[0] = "Apple"
	fl[1] = "Banana"
	fl[2] = "Cucumber"
	fmt.Println(fl)
	fmt.Println("Length of Array :", len(fl))

	// B) Second Type of Array Init
	vl := [3]string{"potato", "beans", "mushroom"}
	fmt.Println(vl)
}
