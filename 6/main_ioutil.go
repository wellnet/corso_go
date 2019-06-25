package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main1() {
	// Read file in memory.
	dat, err := ioutil.ReadFile("file.txt")
	if err != nil {
		panic(err)
	}
	text := string(dat)

	// Print some parts of the file.
	fmt.Println("Old file:")
	fmt.Println(text[:59])           // First 59 bytes.
	fmt.Println(text[len(text)-53:]) // and last 53 bytes.

	// Write data in variable.
	text = strings.Replace(text, "Carmack", "Romero", -1)
	dat2 := []byte(text)
	err = ioutil.WriteFile("file2.txt", dat2, 0644)
	if err != nil {
		panic(err)
	}

	// Print some parts of the new data written to file.
	fmt.Println("New file:")
	fmt.Println(text[:59])           // First 59 bytes.
	fmt.Println(text[len(text)-53:]) // and last 53 bytes.

}
