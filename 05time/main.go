package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study of Golang!")
	presentTime := time.Now()

	fmt.Println(presentTime)

	fmt.Printf("Type of time is: %T \n", presentTime)

	fmt.Println(presentTime.Format("2006-01-02 15:04:05 Monday"))

	createdDate := time.Date(2023, time.June, 12, 0, 0, 0, 0, time.UTC)
	fmt.Println(createdDate)
	fmt.Println(createdDate.Format("2006-01-02 Monday"))

}
