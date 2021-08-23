//https://open.kattis.com/problems/primepath
package main

import (
	"container/list"
	"fmt"
	"strconv"
)

type Current struct {
	Val       int
	IsVisited map[int]bool
	Step      int
}

func isOk(a, b int) bool {
	x := strconv.Itoa(a)
	y := strconv.Itoa(b)
	diff := 0
	for i := 0; i < len(x); i++ {
		if x[i] != y[i] {
			diff++
		}
	}
	if diff == 1 {
		return true
	}
	return false
}

func main() {
	var isPrime [10000]bool
	for i := 0; i < 10000; i++ {
		isPrime[i] = true
	}
	isPrime[1] = false
	for i := 2; i < 10000; i++ {
		if isPrime[i] {
			for j := i + i; j < 10000; j += i {
				isPrime[j] = false
			}
		}
	}
	var hasPath [10000][10000]bool
	var shortest [10000][10000]int
	for i := 1000; i < 10000; i++ {
		for j := 1000; j < 10000; j++ {
			if i != j && isPrime[i] && isPrime[j] && isOk(i, j) {
				hasPath[i][j] = true
				hasPath[j][i] = true
				shortest[i][j] = 10000
				shortest[j][i] = 10000
			}
		}
	}

	var n, a, b int
	fmt.Scanln(&n)

	for i := 0; i < n; i++ {
		fmt.Scanln(&a, &b)
		var isVisited = map[int]bool{}
		isVisited[a] = true
		q := list.New()
		q.PushBack(Current{
			Val:       a,
			IsVisited: isVisited,
			Step:      0,
		})

		var result = -1
		for q.Len() > 0 {
			e := q.Front()
			current := e.Value.(Current)

			if current.Val == b {
				result = current.Step
				break
			} else {
				for x := 1000; x < 10000; x++ {
					if shortest[current.Val][x] >= current.Step+1 && !current.IsVisited[x] && isPrime[x] && hasPath[current.Val][x] {
						var isVisitedNew = current.IsVisited
						isVisitedNew[x] = true
						q.PushBack(Current{
							Val:       x,
							IsVisited: isVisitedNew,
							Step:      current.Step + 1,
						})
					}
				}
			}
			q.Remove(e)
		}
		if result == -1 {
			fmt.Println("Impossible")
		} else {
			fmt.Println(result)
		}
	}
}
