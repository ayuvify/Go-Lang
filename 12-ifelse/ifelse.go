package main

import "fmt"

func main() {

	// If Else
	var result string
	if age := 23; age < 18{
		result = "Young"
	} else if age == 18{
		result = "On Point"
	} else{
		result = "Old"
	}
	fmt.Print(result)
}