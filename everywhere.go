//https://open.kattis.com/problems/everywhere
package main

import "fmt"

func main() {
	var T, n int
	var city string
	fmt.Scanln(&T)

	for t := 0; t < T; t++ {
		var visited = map[string]bool{}
		fmt.Scanln(&n)
		for i := 0; i < n; i++ {
			fmt.Scanln(&city)
			visited[city] = true
		}
		fmt.Println(len(visited))
	}
}
