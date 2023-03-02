package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to classes of slices!")

	var fruitList = []string{"apple", "mango", "guava"}
	fmt.Println("Fruit list is: ", fruitList)
	fmt.Printf("Type of fruitlist %T \n", fruitList)
	fruitList = append(fruitList, "peach", "banans")
	fmt.Println("Fruit list is: ", fruitList)

	//slicing

	fruitList = fruitList[1:5] // range are non inclusive at end
	fmt.Println("Fruit list is: ", fruitList)

	// make

	highScores := make([]int, 4)
	highScores[0] = 452
	highScores[1] = 42
	highScores[2] = 12
	highScores[3] = 45

	fmt.Println("High Scores", highScores)

	highScores = append(highScores, 23, 65)

	fmt.Println("High Scores", highScores)
	fmt.Println("Sorted: ", sort.IntsAreSorted(highScores))
	sort.Ints(highScores)

	fmt.Println("High Scores", highScores)
	fmt.Println("Sorted: ", sort.IntsAreSorted(highScores))

	// removing a value in slice

	var courses = []string{"react", "python", "golang", "rust"}
	fmt.Println("Current courses are: ", courses)

	courses = append(courses[:2], courses[3:]...)

	fmt.Println("Current courses are: ", courses)

}
