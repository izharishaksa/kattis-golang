//https://open.kattis.com/problems/different
package main

import (
	"bufio"
	"fmt"
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
		if a > b {
			fmt.Println(a - b)
		} else {
			fmt.Println(b - a)
		}
	}
}
