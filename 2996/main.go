package main

/**
 * @param {string[]} strs
 * @return {string}

	Input: nums = [1,2,3,2,5]
	Output: 6
	Explanation: The longest sequential prefix of nums is [1,2,3] with a sum of 6. 6 is not in the array,
	therefore 6 is the smallest missing integer greater than or equal to the sum of the longest sequential prefix.

	Input: nums = [3,4,5,1,12,14,13]
	Output: 15
	Explanation: The longest sequential prefix of nums is [3,4,5] with a sum of 12. 12, 13, and 14 belong to the array while 15 does not.
	Therefore 15 is the smallest missing integer greater than or equal to the sum of the longest sequential prefix.


**/

func missingInteger(nums []int) int {
	if len(nums) == 0 {
		return 1
	}
	sum := 0
	for i := 0; i < len(nums); i++ {
		if i == 0 || nums[i] == nums[i-1]+1 {
			sum += nums[i]
		} else {
			break
		}
	}
	for i := sum; ; i++ {
		found := false
		// [0->3, 1->4, 2->5, 3->1, 4->12, 5->14, 6->13]
		for _, v := range nums {
			if v == i {
				found = true
				break
			}
		}
		if !found {
			return i
		}
	}
}

func main() {
	test := missingInteger([]int{3, 4, 5, 1, 12, 14, 13})
	println(test)
}
