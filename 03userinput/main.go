package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to User Input Program!"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating of pizza:")

	// comma ok syntax // err ok

	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating, ", input)
	fmt.Printf("Type of this rating is %T \n", input)
}
