package main

import "fmt"

func main() {

	var arr [2]int // array of 2 integers
	intArr := [10]int{1, 2, 3, 4, 5} // array of 10 integers with first 5 elements initialized
	inArr2 := [...]int{1, 2, 3, 4, 5} // array of 5 integers with size inferred from the number of elements


	arr[0] = 1
	arr[1] = 2
	fmt.Println(arr)
	fmt.Println(intArr)
	fmt.Println(inArr2)

}