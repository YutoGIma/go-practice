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
	tableNum := strings.Split(sc.Text(), " ")
	row, _ := strconv.Atoi(tableNum[0])
	column, _ := strconv.Atoi(tableNum[1])
	var table [][]string
	var miniTable []string

	for i := 0; i < row; i++ {
		sc.Scan()
		data := sc.Text()
		miniTable = strings.Split(data, "")
		table = append(table, miniTable)
	}
	var rowBomb bool = false
	var rowBombNumber []int
	for n := 0; n < row; n++ {
		if arraystrCheck(table[n], "#") {
			rowBomb = true
		}
		for m := 0; m < column; m++ {
			if table[n][m] == "#" {
				rowBombNumber = append(rowBombNumber, m)
			}
			if rowBomb {
				table[n][m] = "#"
			}
		}
		rowBomb = false
	}

	var count int

	for n := 0; n < row; n++ {
		for m := 0; m < column; m++ {
			if arrayCheck(rowBombNumber, m) {
				table[n][m] = "#"
			}
			if table[n][m] == "#" {
				count = count + 1
			}
		}
	}

	fmt.Print(count)

}

func arrayCheck(array []int, num int) bool {
	for _, v := range array {
		if v == num {
			return true
		}
	}
	return false
}

func arraystrCheck(array []string, num string) bool {
	for _, v := range array {
		if v == num {
			return true
		}
	}
	return false
}
