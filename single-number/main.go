package main

import "fmt"

/*
l = 0 ^ 4 = 4
l = 4 ^ 1 = 5
l = 5 ^ 2 = 7
l = 7 ^ 1 = 6
l = 6 ^ 2 = 4
*/
func singleNumber(nums []int) int {
	l := 0
	for i := 0; i < len(nums); i++ {
		l ^= nums[i]
	}
	return l
}

func main() {
	nums := []int{4, 1, 2, 1, 2}
	fmt.Println(singleNumber(nums))
}
