//https://open.kattis.com/problems/baloni
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var N, H, total int
	fmt.Fscan(reader, &N)
	flag := map[int]int{}
	for i := 0; i < N; i++ {
		fmt.Fscan(reader, &H)
		if val, ok := flag[H+1]; ok && val > 0 {
			flag[H+1]--
			flag[H]++
		} else {
			flag[H]++
			total++
		}
	}
	fmt.Println(total)
}
