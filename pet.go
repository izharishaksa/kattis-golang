//https://open.kattis.com/problems/pet
package main

import "fmt"

func main() {
	var winner, max int
	for i := 0; i < 5; i++ {
		var curScore, score int
		for j := 0; j < 4; j++ {
			fmt.Scanf("%d", &score)
			curScore += score
		}
		if curScore > max {
			max = curScore
			winner = i + 1
		}
	}
	fmt.Println(winner, max)
}
