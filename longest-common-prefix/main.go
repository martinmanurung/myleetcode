package main

import "fmt"

/**
 * @param {string[]} strs
 * @return {string}

 Buat sebuah fungsi untuk menemukan string awalan (prefix) terpanjang yang sama di antara sekumpulan string

 Example 1:

	Input: strs = ["flower","flow","flight"]
	Output: "fl"


Example 2:

	Input: strs = ["dog","racecar","car"]
	Output: ""
	Explanation: There is no common prefix among the input strings.

	i < len(strs[0]) -> untuk iterasi setiap karakter dari string pertama
	char := strs[0][i] -> untuk menyimpan karakter saat ini dari string pertama

	j < len(strs) -> untuk iterasi setiap string dalam array

	i >= len(strs[j]) -> untuk memeriksa apakah indeks karakter saat ini melebihi panjang string saat ini

	strs[j][i] != char -> untuk memeriksa apakah karakter saat ini dari string saat ini tidak sama dengan karakter dari string pertama

	return strs[0][:i] -> untuk mengembalikan awalan terpanjang yang sama di antara semua string

	f

*/

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	for i := 0; i < len(strs[0]); i++ {
		char := strs[0][i]
		for j := 1; j < len(strs); j++ {
			if i >= len(strs[j]) || strs[j][i] != char {
				return strs[0][:i]
			}
		}
	}

	return strs[0]
}

func main() {
	test := longestCommonPrefix([]string{"flower", "flower", "flower", "flower"})
	fmt.Println(test)
}
