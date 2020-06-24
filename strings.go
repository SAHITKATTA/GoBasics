package main

import (
	"fmt"
)

func main() {
	book := "the color of magic"
	fmt.Println(book)
	fmt.Println(len(book))
	fmt.Printf("book[0]=%v (Type is %T) ", book[0], book[0])
	// book[0] = 116 ;String in go are immutable

	// Slicing
	fmt.Println(book[4:11])
	// Concatenation
	fmt.Println("t" + book[1:])
	//unicode support
	//multiline
	poem := `
	the road goes ever on
	Down from the dorr where it began
	`
	fmt.Println(poem)
}
