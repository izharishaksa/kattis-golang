//https://open.kattis.com/problems/stopwatch
package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scanln(&n)
	var pressed = make([]int32, n)
	for i := 0; i < n; i++ {
		fmt.Scanln(&pressed[i])
	}
	if n%2 == 1 {
		fmt.Println("still running")
	} else {
		var result int32
		for i := len(pressed) - 1; i >= 0; i -= 2 {
			result += pressed[i] - pressed[i-1]
		}
		fmt.Println(result)
	}
}
