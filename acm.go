//https://open.kattis.com/problems/acm
package main

import (
	"fmt"
)

func main() {
	var minute int
	var problemId string
	var status string
	var isSolved map[string]bool
	var penalty map[string]int
	var totalTime int
	var totalSolve int

	isSolved = map[string]bool{}
	penalty = map[string]int{}
	for {
		fmt.Scanf("%d", &minute)
		if minute == -1 {
			break
		}
		fmt.Scanf("%s %s", &problemId, &status)
		if status == "right" && !isSolved[problemId] {
			totalSolve++
			totalTime += minute + penalty[problemId]
			isSolved[problemId] = true
		} else {
			penalty[problemId] += 20
		}
	}
	fmt.Println(totalSolve, totalTime)
}
