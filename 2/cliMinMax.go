package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Give me more numbers.")
		os.Exit(1)
	}

	// Read os.Args: [0]=program, [1]=number, [2]=number, ..., [n]=number
	arguments := os.Args
	min, err := strconv.ParseFloat(arguments[1], 64)
	if err != nil {
		fmt.Println("Must be a number!")
		os.Exit(1)
	}
	max := min

	// Find the min-max.
	for i := 2; i < len(arguments); i++ {
		n, err := strconv.ParseFloat(arguments[i], 64)
		if err != nil {
			fmt.Println("The argument!")
			os.Exit(1)
		}
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}

	fmt.Println("Min:", min)
	fmt.Println("Max:", max)
}
