package main

import "fmt"

// ini kalau pakai string, tapi kalau pakai int, bisa langsung pakai math
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x < 10 {
		return true
	}
	s := fmt.Sprintf("%d", x)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}

// ini kalau gaboleh ada operasi string
func isPalindromeInt(x int) bool {
	if x < 0 {
		return false
	}
	reversed := 0
	original := x
	for x != 0 {
		reversed = reversed*10 + x%10
		x /= 10
	}
	return reversed == original
}

// 121 -> "121" -> len(2) -> 0 < 2 ->

func main() {
	fmt.Println(isPalindrome(121))      // true
	fmt.Println(isPalindrome(-121))     // false
	fmt.Println(isPalindrome(10))       // false
	fmt.Println(isPalindromeInt(12321)) // true
}
