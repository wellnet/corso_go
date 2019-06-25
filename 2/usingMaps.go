package main

import (
	"fmt"
)

func main() {

	// Create.
	myMap := make(map[string]int)
	myMap["key1"] = 21
	myMap["key2"] = 42
	myMap["key3"] = 84
	fmt.Println("myMap:", myMap)

	// Delete.
	anotherMap := map[string]int{
		"key1": 21,
		"key2": 42,
	}

	fmt.Println("anotherMap:", anotherMap)
	delete(anotherMap, "key1")
	delete(anotherMap, "key1")
	delete(anotherMap, "key1")
	fmt.Println("anotherMap:", anotherMap)

	// Check if the key exits.
	_, ok := myMap["newKey"]
	if ok {
		fmt.Println("The key exists!")
	} else {
		fmt.Println("The key does NOT exist")
	}

	// Range over map.
	for key, value := range myMap {
		fmt.Println(key, value)
	}
}
