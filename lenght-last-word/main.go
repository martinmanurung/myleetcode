package main

import (
	"fmt"
	"strings"
)

/*
if the last word is empty, use the previous word
*/
func lengthOfLastWord(s string) int {
	s = strings.TrimSpace(s)
	st := strings.Split(s, " ")
	word := st[len(st)-1]
	return len(word)
}

func main() {
	fmt.Println(lengthOfLastWord("   fly me   to   the moon  "))
}
