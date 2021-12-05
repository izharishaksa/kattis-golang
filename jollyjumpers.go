//https://open.kattis.com/problems/jollyjumpers
package main

import (
	"fmt"
	"math"
)

func main() {
	var N int
	for {
		_, err := fmt.Scan(&N)
		if err != nil {
			break
		}
		var arr = make([]int, N)
		for i := 0; i < N; i++ {
			fmt.Scan(&arr[i])
		}
		flag := map[int]bool{}
		isJolly := true
		for i := 1; i < N; i++ {
			diff := int(math.Abs(float64(arr[i] - arr[i-1])))
			if diff > N-1 {
				isJolly = false
				break
			}
			flag[diff] = true
		}
		if len(flag) == N-1 && isJolly {
			fmt.Println("Jolly")
		} else {
			fmt.Println("Not jolly")
		}
	}
}
