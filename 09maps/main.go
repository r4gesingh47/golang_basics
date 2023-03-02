package main

import "fmt"

func main() {
	fmt.Println("Welcome to classes of maps!")
	languages := make(map[string]string)

	languages["JS"] = "JavaScript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	fmt.Println("List of all languages: ", languages)
	fmt.Println("JS shorts for : ", languages["JS"])

	delete(languages, "RB")

	fmt.Println("List of all languages: ", languages)

	// loops are interesting in golang

	for key, value := range languages {
		fmt.Printf("For key: %v, value: %v\n", key, value)
	}
}
