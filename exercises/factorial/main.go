package main

import "fmt"

func main() {

	fmt.Println("Enter A Number")
	var n int
	_, _ = fmt.Scan(&n)
	dx := fact(n)
	fmt.Println(dx)
}

func fact(n int) int {
	res := 1
	for i := 1; i <= n; i++ {
		res = res * i
	}
	return res
}
