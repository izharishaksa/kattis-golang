package main

import (
	"bufio"
	"fmt"
	"os"
)

type Data struct {
	Val int64
	Len int
	seq []int64
}

func (d *Data) appendToSeq(number []int64) {
	d.seq = append(number, d.seq...)
}

func compareSeq(current []int64, candidate []int64) bool {
	//fmt.Println(candidate, current)
	for i := 0; i < len(current); i++ {
		if candidate[i] < current[i] {
			return true
		}
	}
	return false
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	var N int
	var number int64
	for {
		fmt.Fscan(reader, &N)
		if N == 0 {
			break
		}
		arr := make([]Data, N)
		for i := 0; i < N; i++ {
			fmt.Fscan(reader, &number)
			arr[i] = Data{
				Val: number,
				Len: 1,
			}
			arr[i].appendToSeq([]int64{number})
		}

		maxLen := 1
		seq := []int64{arr[0].Val}
		for i := 1; i < N; i++ {
			max := 1
			var candidate []int64
			for j := i - 1; j >= 0; j-- {
				if arr[i].Val > arr[j].Val && arr[i].Len+arr[j].Len > max {
					max = arr[i].Len + arr[j].Len
					candidate = arr[j].seq
				} else if arr[i].Val > arr[j].Val && arr[i].Len+arr[j].Len == max && (len(candidate) == 0 || compareSeq(candidate, arr[j].seq)) {
					max = arr[i].Len + arr[j].Len
					candidate = arr[j].seq
				}
			}
			arr[i].Len = max
			arr[i].appendToSeq(candidate)

			//fmt.Println(arr[i].Val, maxLen, arr[i].Len, arr[i].seq, seq, compareSeq(seq, arr[i].seq))
			if arr[i].Len > maxLen || (maxLen == arr[i].Len && compareSeq(seq, arr[i].seq)) {
				maxLen = arr[i].Len
				seq = arr[i].seq
			}
		}
		//fmt.Println(arr)
		fmt.Print(maxLen, " ")
		for i := 0; i < len(seq); i++ {
			fmt.Print(seq[i], " ")
		}
		fmt.Println()
	}
}