package main

import (
	"fmt"
)

func mostPoints(questions [][]int) int64 {
	n := len(questions)
	ans := make([]int, n+1)

	for i := n - 1; i >= 0; i-- {
		point := questions[i][0]
		brainpower := questions[i][1]
		skipId := i + brainpower + 1
		take := point
		if skipId < n {
			take += ans[skipId]
		}
		skip := ans[i+1]
		ans[i] = max(take, skip)
	}
	return int64(ans[0])
}

func main() {
	fmt.Println(mostPoints([][]int{
		{3, 2},
		{4, 3},
		{4, 4},
		{2, 5}}))
}
