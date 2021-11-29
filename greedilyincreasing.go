package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	buffer := new(bytes.Buffer)
	defer writer.Flush()

	var N, leftMost, current, total int

	fmt.Fscan(reader, &N)
	fmt.Fscan(reader, &leftMost)
	buffer.WriteString(strconv.Itoa(leftMost))
	buffer.WriteString(" ")
	total++
	for i := 1; i < N; i++ {
		fmt.Fscan(reader, &current)
		if current > leftMost {
			leftMost = current
			buffer.WriteString(strconv.Itoa(leftMost))
			buffer.WriteString(" ")
			total++
		}
	}
	fmt.Fprintln(writer, total)
	fmt.Fprintln(writer, buffer.String())
}
