package main

import (
	"fmt"
)

func main() {
	x := 10
	switch x {
	case 1:
		fmt.Println("x is one")
	case 2:
		fmt.Println("x is two")
	case 3:
		fmt.Println("x is three")
	default:
		fmt.Println("many")
	}
	switch {
	case x > 100:
		fmt.Println("x is greater than 100")
	case x < 100:
		fmt.Println("x is less than 100")
	}
}
