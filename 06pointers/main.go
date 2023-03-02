package main

import "fmt"

func main() {
	fmt.Println("Welcome to class for pointers!")

	var ptr *int

	fmt.Println("Value of pointer is :", ptr)
	fmt.Printf("Type of pointer is %T \n", ptr)

	myNumber := 12
	fmt.Println("Saying ptr to point to myNumber reference")
	ptr = &myNumber // Referencing
	fmt.Println("Value of pointer is :", ptr)
	fmt.Println("Accesing the value using ptr: ", *ptr)

	// changingt the value of myNumber using pointer

	*ptr = *ptr * 2
	fmt.Println("New Values is: ", myNumber)
}
