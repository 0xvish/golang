package main

import (
	"fmt"
	"strings"
)

func main() {

	// Strings, Runes and Bytes in Go
	// Strings in Go are made up of runes, which are Unicode code points.
	// A rune is a character, and a string is a sequence of characters.
	// Strings are made up of bytes, which are 8-bit integers that represent ASCII characters.
	// In Go, a string is in effect a read-only slice of bytes.

	var myString string = "⚡Vish⚡" //String with UTF-8 characters
	var indexed = myString[0]      //Accessing the first byte of the string
	fmt.Println("First byte of the string: ", indexed)
	for i, v := range myString {
		fmt.Println(i, v) //Notice the indexing here is not contiguous, it is because some UTF-8 need multiple bytes to represent a character
	}

	//To address the above issue, we can convert the string to a slice of runes
	var myRunes = []rune(myString)
	for i, v := range myRunes {
		fmt.Println(i, v) //Notice the indexing here is contiguous but the values are runes, thus we see the Unicode code points. Multiple bytes are used to represent a single character
	}

	//To address the above issue, we can convert the string to a slice of bytes
	var myBytes = []byte(myString)
	for i, v := range myBytes {
		fmt.Println(i, v) //Notice the indexing here is contiguous and the values are bytes, thus we see the ASCII values
	}

	//To convert a rune to a string
	var myRune rune = '⚡'
	var myRuneString = string(myRune)
	fmt.Println("Rune: ", myRune)
	fmt.Println("String: ", myRuneString)


	var strSlice = []string{"Vish", "is", "learning", "Go"}
	var str = ""
	for _, v := range strSlice {
		str += v + " "
	}
	fmt.Println(str) //string are immutable in Go, and this method results in creating a new string every time we concatenate a string

	//The best way to concatenate strings in Go is to use the strings package, because it uses a buffer to store the strings and is more efficient
	var strSlice2 = []string{"Vish", "is", "learning", "Go"}
	var strBuilder strings.Builder
	for _, v := range strSlice2 {
		strBuilder.WriteString(v)
		strBuilder.WriteString(" ")
	}
	fmt.Println(strBuilder.String())

}