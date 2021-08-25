//https://open.kattis.com/problems/missingnumbers
package main

import "fmt"

func main() {
	var n, last int
	var isCounted [201]bool
	fmt.Scanln(&n)
	for i := 0; i < n; i++ {
		fmt.Scanln(&last)
		isCounted[last] = true
	}
	var count = 0
	for i := 1; i < last; i++ {
		if !isCounted[i] {
			fmt.Println(i)
			count++
		}
	}
	if count == 0 {
		fmt.Println("good job")
	}

}
