package main

import "fmt"

/**
 * Input: nums = [0,0,1,1,1,2,2,3,3,4]
 * Output: 5, nums = [0,1,2,3,4,_,_,_,_,_]
 * Explanation: Your function should return k = 5, with the first five elements of nums being 0, 1, 2, 3, and 4 respectively.
 *
 *
 * 	Input: nums = [1,1,2]
 * 	Output: 2, nums = [1,2,_]
 * 	Explanation: Your function should return k = 2, with the first two elements of nums being 1 and 2 respectively.
 */

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// k adalah indeks untuk menempatkan elemen unik berikutnya
	k := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			nums[k] = nums[i]
			k++
		}
	}
	fmt.Println(nums)

	return k
}

func main() {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Println(removeDuplicates(nums))
}
