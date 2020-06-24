package main

import (
	"fmt"
)

func main() {
	loons := []string{"bugs", "daffy", "taz"}
	fmt.Printf("loons = %v (type %T)\n", loons, loons)
	fmt.Println(len(loons))

	fmt.Println("------------")
	fmt.Println(loons[1])
	fmt.Println(loons[1:])

	fmt.Println("------------")
	for i := range loons {
		fmt.Println(i)
	}
	for i, name := range loons {
		fmt.Printf("%s at %d\n", name, i)
	}
	//double value reange if ignore the index then
	for _, name := range loons {
		fmt.Println(name)
	}
	loons = append(loons, "elmer")
	for _, name := range loons {
		fmt.Println(name)
	}
}
