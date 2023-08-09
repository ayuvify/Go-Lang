package main

import (
	"fmt"
	"sort"
)

func main() {

	// Sort Function
	a := make([]int, 4)
	sort.Ints((a))
	fmt.Println(a)


	// Sort Check Function
	b := make([]int, 4)
	fmt.Println(sort.IntsAreSorted(b))


	// Example of Sort Function
	var course = []string{"react","js","swift","python","ruby"}
	fmt.Println(course)
	index := 2
	course = append(course[:index], course[index+1:]...)
	fmt.Println(course)
}