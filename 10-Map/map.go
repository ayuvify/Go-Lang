package main

import "fmt"

func main() {

	// Maps
	lang := make(map[string]string)

	lang["JS"] = "Javascript"
	lang["RB"] = "Ruby"
	lang["PY"] = "Python"
	fmt.Println("List of all languages: ", lang)
	fmt.Println("JS Full Form: ", lang["JS"])

	// Delete Element From Map
	delete(lang, "RB")
	fmt.Println("List of all languages: ", lang)

	// loops on maps
	for key, value := range lang {
		fmt.Printf("For key %v, value is %v\n", key, value)
	}

	for _, value := range lang {
		fmt.Printf("For value is %v\n", value)
	}
}
