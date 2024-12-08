package main

import "fmt"

func main() {

	var slice []int // slice of integers
	intSlice := []int{1, 2, 3, 4, 5} // slice of integers with first 5 elements initialized

	slice = append(slice, 1)
	slice = append(slice, 2)
	fmt.Println(slice)
	fmt.Println("Length of slice: ", len(slice))
	fmt.Println(intSlice)
	fmt.Println("Length of intSlice: ", len(intSlice))

	// Slicing a slice
	fmt.Println(intSlice[1:3]) // [2 3]
	fmt.Println(intSlice[:3]) // [1 2 3]
	fmt.Println(intSlice[3:]) // [4 5]

	// Iterating over a slice
	for i, v := range intSlice {
		fmt.Println(i, v)
	}

	// Iterating over a slice without using index
	for _, v := range intSlice {
		fmt.Println(v)
	}

	// Iterating over a slice without using value
	for i := range intSlice {
		fmt.Println(intSlice[i])
	}


}