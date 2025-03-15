package main

import "fmt"

func main() {
	res := fac(1, 50)
	fmt.Print(res)
}

func fac(low, high int) []int {
	var arr []int
	for i := low; i <= high; i++ {
		if i%7 == 0 && i%5 != 0 {
			arr = append(arr, i)
		}
	}
	return arr
}
