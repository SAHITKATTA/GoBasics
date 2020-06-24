//first and last digit should be same when a 4 digit is multiplied to other numbers

package main

import (
	"fmt"
)

func main() {
	count := 0
	for i := 1000; i < 10000; i++ {
		for j := i; j < 10000; j++ {
			s := fmt.Sprint(i * j)
			// s:= fmt.Sprintf("%d",i*j)
			if s[0] == s[len(s)-1] {
				count++
			}
		}
	}
	fmt.Println("the count is", count)
}
