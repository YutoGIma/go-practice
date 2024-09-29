package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	rowCount, _ := strconv.Atoi(sc.Text())
	columnCount := 5

	var table [][]string
	var row []string

	for i := 0; i < rowCount; i++ {
		sc.Scan()
		data := sc.Text()
		row = strings.Split(data, "")
		table = append(table, row)
	}

	for {
		deleteCenter := checkCross(table, rowCount, columnCount)
		table = changeTable(table, deleteCenter, rowCount, columnCount)
		table = dropNumber(table, rowCount, columnCount)
		if len(deleteCenter) == 0 {
			break
		}
		deleteCenter = nil
	}

	for _, row := range table {
		for _, point := range row {
			fmt.Print(point)
		}
		fmt.Print("\n")
	}

}

func checkCross(table [][]string, rowCount int, columnCount int) [][]int {
	var deleteCenter [][]int
	for h := 0; h < rowCount; h++ {
		for w := 0; w < columnCount; w++ {
			if table[h][w] == "." {
				continue
			}
			if h == 0 {
				if w == 0 {
					if table[0][0] == table[0][1] && table[0][0] == table[1][0] {
						deleteCenter = append(deleteCenter, []int{h, w})
					}
				} else if w == (columnCount - 1) {
					if table[0][w] == table[0][(w-1)] && table[0][w] == table[1][w] {
						deleteCenter = append(deleteCenter, []int{h, w})
					}
				} else {
					if table[0][w] == table[1][w] && table[0][w] == table[0][(w+1)] && table[0][w] == table[0][(w-1)] {
						deleteCenter = append(deleteCenter, []int{h, w})
					}
				}
			} else if h == (rowCount - 1) {
				if w == 0 {
					if table[h][0] == table[h][1] && table[h][0] == table[(h - 1)][0] {
						deleteCenter = append(deleteCenter, []int{h, w})
					}
				} else if w == (columnCount - 1) {
					if table[h][w] == table[h][(w-1)] && table[h][w] == table[(h - 1)][w] {
						deleteCenter = append(deleteCenter, []int{h, w})
					}
				} else {
					if table[h][w] == table[(h - 1)][w] && table[h][w] == table[h][(w+1)] && table[h][w] == table[h][(w-1)] {
						deleteCenter = append(deleteCenter, []int{h, w})
					}
				}
			} else {
				if w == 0 {
					if table[h][0] == table[h][1] && table[h][0] == table[(h + 1)][0] && table[h][0] == table[(h - 1)][0] {
						deleteCenter = append(deleteCenter, []int{h, w})
					}
				} else if w == (columnCount - 1) {
					if table[h][w] == table[h][(w-1)] && table[h][w] == table[(h + 1)][w] && table[h][w] == table[(h - 1)][w] {
						deleteCenter = append(deleteCenter, []int{h, w})
					}
				} else {
					if table[h][w] == table[(h + 1)][w] && table[h][w] == table[(h - 1)][w] && table[h][w] == table[h][(w+1)] && table[h][w] == table[h][(w-1)] {
						deleteCenter = append(deleteCenter, []int{h, w})
					}
				}
			}
		}
	}
	return deleteCenter
}

func changeTable(table [][]string, deleteCenter [][]int, rowCount int, columnCount int) [][]string {
	for _, point := range deleteCenter {
		table[point[0]][point[1]] = "."
		if point[0] != 0 {
			table[(point[0] - 1)][point[1]] = "."
		}
		if point[1] != 0 {
			table[point[0]][(point[1] - 1)] = "."
		}
		if point[0] != (rowCount - 1) {
			table[(point[0] + 1)][point[1]] = "."
		}
		if point[1] != (columnCount - 1) {
			table[point[0]][(point[1] + 1)] = "."
		}
	}
	return table
}

func dropNumber(table [][]string, rowCount int, columnCount int) [][]string {
	var rowNumber []string
	var tmp []string
	for c := 0; c < columnCount; c++ {
		rowNumber = nil
		for r := (rowCount - 1); r >= 0; r-- {
			if table[r][c] != "." {
				rowNumber = append(rowNumber, table[r][c])
			}
		}
		for r := (rowCount - 1); r >= 0; r-- {
			if len(rowNumber) > 0 {
				table[r][c] = rowNumber[0]
				rowNumber = append(tmp, rowNumber[1:]...)
				tmp = nil
			} else {
				table[r][c] = "."
			}
		}
	}
	return table
}
