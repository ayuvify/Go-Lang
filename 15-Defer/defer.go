package main

import "fmt"

func main() {

	// It is a defer statement[Function]
	defer func() {
		for j := 0; j < 15; {
			if j%2 == 0 {
				fmt.Println(j)
			}
			j++
		}
	}()

	// It is a defer statement[Print]
	defer fmt.Println("Done")

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

}
