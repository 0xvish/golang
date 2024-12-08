package main

import "fmt"

func main() {

	var myMap map[string]int // map with string keys and integer values
	myMap = make(map[string]int)
	myMap["one"] = 1
	myMap["two"] = 2
	myMap["three"] = 3
	fmt.Println(myMap)
	fmt.Println("Length of myMap: ", len(myMap))

	// Iterating over a map
	for k, v := range myMap {
		fmt.Println(k, v)
	}

	// Iterating over a map without using key
	for _, v := range myMap {
		fmt.Println(v)
	}

	// Iterating over a map using key
	for k := range myMap {
		fmt.Println(myMap[k])
	}

	// Accessing a value from a map
	var value, ok = myMap["one"]
	if ok {
		fmt.Println("Value of key 'one': ", value)
	} else {
		fmt.Println("Key 'one' not found")
	}

}