package main

import (
	"io/ioutil"
	"strings"
)

func main() {
	// leggere il file
	data, err := ioutil.ReadFile("file.txt")
	if err != nil {
		panic(err)
	}
	text := string(data)

	// cambiare qualcosa
	text = strings.Replace(text, "Carmack", "Romero", -1)

	// scrivere su un nuovo
	err = ioutil.WriteFile("file2.txt", []byte(text), 0644)
	if err != nil {
		panic(err)
	}

}
