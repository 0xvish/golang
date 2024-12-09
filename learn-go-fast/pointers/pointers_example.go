package main

import "fmt"

func main() {
	var p *int = new(int)
	fmt.Println("p: ", p) // address of 0, default value of int
	fmt.Println("&p: ", &p) // address of p
	fmt.Println("*p: ", *p) // 0

	var x int = 5
	p = &x
	fmt.Println("p: ", p) // address of x
	fmt.Println("&p: ", &p) // address of p
	fmt.Println("*p: ", *p) // 5

}