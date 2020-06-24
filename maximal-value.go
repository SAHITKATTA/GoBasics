package main

import "fmt"

func main() {
	nums := []int{16, 8, 42, 4, 23, 15}
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if max < nums[i] {
			max = nums[i]
		}
	}
	fmt.Println(max)
}
