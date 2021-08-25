//https://open.kattis.com/problems/different
package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := strings.Split(scanner.Text(), " ")
		a, _ := strconv.ParseUint(input[0], 10, 64)
		b, _ := strconv.ParseUint(input[1], 10, 64)
		fmt.Println(int64(math.Abs(float64(a) - float64(b))))
	}
}
