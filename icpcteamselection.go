//https://open.kattis.com/problems/icpcteamselection
package main

import (
	"fmt"
	"sort"
)

func main() {
	var T, N int
	fmt.Scanf("%d", &T)
	for t := 0; t < T; t++ {
		fmt.Scanf("%d", &N)
		var arr = make([]int, N*3)
		index := 0
		for i := 0; i < N; i++ {
			for j := 0; j < 3; j++ {
				fmt.Scanf("%d", &arr[index])
				index++
			}
		}
		sort.Ints(arr)
		total := 0
		for i := N*3 - 2; N > 0; i -= 2 {
			total += arr[i]
			N--
		}
		fmt.Println(total)
	}
}
