package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcom to array class!")

	var frutiList [4]string

	frutiList[0] = "Apple"
	frutiList[1] = "Guava"
	frutiList[3] = "Mango"

	fmt.Println("Frutlist is: ", frutiList)
	fmt.Println("Length of the fruit list is: ", len(frutiList))

	vegetableList := [3]string{"potato", "beans", "mushroom"}

	fmt.Println("Vegetable list is: ", vegetableList)

	arr := [...]string{"a", "b", "c"}       // let the compiler determine the length at runtime
	fmt.Printf("Distinctive Type: %T", arr) // [3]string is a distinctive type rather than [5] string
	fmt.Println(arr)

}
