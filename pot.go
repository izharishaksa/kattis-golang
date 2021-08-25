package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	var n, total int
	var x, y int64
	var str, a, b string
	fmt.Scanln(&n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%s", &str)
		a = str[:len(str)-1]
		b = str[len(str)-1:]
		x, _ = strconv.ParseInt(a, 10, 64)
		y, _ = strconv.ParseInt(b, 10, 64)
		total += int(math.Pow(float64(x), float64(y)))
	}
	fmt.Println(total)
}
