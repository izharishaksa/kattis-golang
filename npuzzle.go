//https://open.kattis.com/problems/npuzzle
package main

import (
	"fmt"
)

type Position struct {
	Row int
	Col int
}

func main() {
	pos := map[int32]Position{}
	pos['.'] = Position{
		Row: 3,
		Col: 3,
	}
	var row, col int
	for i := 'A'; i < 'P'; i++ {
		pos[i] = Position{
			Row: row,
			Col: col,
		}
		col++
		if col%4 == 0 {
			row++
			col = 0
		}
	}
	var cur string
	var curRow, curCol, scatter int
	for i := 0; i < 4; i++ {
		fmt.Scanln(&cur)
		for j := 0; j < 4; j++ {
			if cur[j] == '.' {
				curCol++
				if curCol%4 == 0 {
					curCol = 0
					curRow++
				}
				continue
			}
			var expPos = pos[int32(cur[j])]
			if expPos.Row != curRow || expPos.Col != curCol {
				if curRow > expPos.Row {
					scatter += curRow - expPos.Row
				} else {
					scatter += expPos.Row - curRow
				}
				if curCol > expPos.Col {
					scatter += curCol - expPos.Col
				} else {
					scatter += expPos.Col - curCol
				}
			}

			curCol++
			if curCol%4 == 0 {
				curCol = 0
				curRow++
			}
		}
	}
	fmt.Println(scatter)
}
