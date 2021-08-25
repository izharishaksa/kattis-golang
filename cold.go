//https://open.kattis.com/problems/cold
package main

import "fmt"

func main() {
	var n, total, degree int
	fmt.Scanln(&n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &degree)
		if degree < 0 {
			total++
		}
	}
	fmt.Println(total)

}
