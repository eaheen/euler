package p8

import "fmt"

// Returns solution for Problem 8
func Sol() {
	fmt.Println(productInSeries(13, "ans/p8/p8.num"))
}

// Interactive solution session for Problem 8
func Ask() {
	fmt.Println("Find the n consecutive digits within x that have the greeatest product.")
	fmt.Print("Input n: ")
	var n int
	fmt.Scanln(&n)
	fmt.Print("Input path to x, stored as ASCII: ")
	var x string
	fmt.Scanln(&x)
	fmt.Println("Answer:", productInSeries(n, x))
}
