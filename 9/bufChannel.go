package main

import (
	"fmt"
)

func main() {
	numbers := make(chan int, 5)
	counter := 10

	for i := 0; i < counter; i++ {
		select {
		case numbers <- i:
			fmt.Println("Insert ", i)
		default:
			fmt.Println("Not enough space for", i)
		}
	}

	for i := 0; i < counter+5; i++ {
		select {
		case num := <-numbers:
			fmt.Println("Extracted ", i)
		default:
			fmt.Println("Nothing more to be done!")
			break
		}
	}
}
