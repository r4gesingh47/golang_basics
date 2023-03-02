package main

import "fmt"

const LoginToken string = "asdljav;oi1342o3432" // this is public attribute so the starting lettter cases decides that

func main() {
	var username string = "Ravi Kumar Singh"
	fmt.Println(username)
	fmt.Printf("Variable is of type : %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type : %T \n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type : %T \n", smallVal)

	var smallFloat float64 = 255.323212121231231231
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type : %T \n", smallFloat)

	// default values and some aliases

	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type : %T \n", anotherVariable)

	// implicit type

	var website = "r4ge.com"
	fmt.Println(website)

	// no var style

	numberOfUser := 300000
	fmt.Println(numberOfUser) // not allowed outside functions

	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type : %T \n", LoginToken)
}
