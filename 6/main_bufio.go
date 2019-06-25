package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func main2() {
	// Read file in memory.
	fi, err := os.Open("file.txt")
	if err != nil {
		panic(err)
	}
	defer fi.Close()
	r := bufio.NewReader(fi)
	content, _ := ioutil.ReadAll(r)

	text := string(content)

	// Print some parts of the file.
	fmt.Println("Old file:")
	fmt.Println(text[:59])           // First 59 bytes.
	fmt.Println(text[len(text)-53:]) // and last 53 bytes.

	// Write data in variable.
	text = strings.Replace(text, "Carmack", "Romero", -1)

	f2, err := os.Create("file2.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = io.WriteString(f2, text)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Print some parts of the new data written to file.
	fmt.Println("New file:")
	fmt.Println(text[:59])           // First 59 bytes.
	fmt.Println(text[len(text)-53:]) // and last 53 bytes.

}
