// HWのターブルが渡されてSから.のとこだけを通って外に出られるかのコード
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
	height, _ := strconv.Atoi(tableNum[0])
	width, _ := strconv.Atoi(tableNum[1])

	var table [][]string
	var row []string
	var start []int

	for i := 0; i < height; i++ {
		sc.Scan()
		data := sc.Text()
		row = strings.Split(data, "")
		table = append(table, row)
	}

	for h := 0; h < height; h++ {
		for w := 0; w < width; w++ {
			if table[h][w] == "S" {
				start = []int{h, w}
			}
		}
	}
	currentPoint := start
	var gonePoint [][]int
	var prevPoint [][]int

	for {
		if currentPoint[0] == 0 || currentPoint[0] == height-1 {
			fmt.Println("YES")
			return
		} else if currentPoint[1] == 0 || currentPoint[1] == width-1 {
			fmt.Println("YES")
			return
		} else {
			if table[currentPoint[0]+1][currentPoint[1]] == "." && notSliceSearch(gonePoint, []int{currentPoint[0] + 1, currentPoint[1]}) && notSliceSearch(prevPoint, []int{currentPoint[0] + 1, currentPoint[1]}) {
				prevPoint = append(prevPoint, currentPoint)
				currentPoint = []int{currentPoint[0] + 1, currentPoint[1]}
				continue
			} else if table[currentPoint[0]][currentPoint[1]+1] == "." && notSliceSearch(gonePoint, []int{currentPoint[0], currentPoint[1] + 1}) && notSliceSearch(prevPoint, []int{currentPoint[0], currentPoint[1] + 1}) {
				prevPoint = append(prevPoint, currentPoint)
				currentPoint = []int{currentPoint[0], currentPoint[1] + 1}
				continue
			} else if table[currentPoint[0]-1][currentPoint[1]] == "." && notSliceSearch(gonePoint, []int{currentPoint[0] - 1, currentPoint[1]}) && notSliceSearch(prevPoint, []int{currentPoint[0] - 1, currentPoint[1]}) {
				prevPoint = append(prevPoint, currentPoint)
				currentPoint = []int{currentPoint[0] - 1, currentPoint[1]}
				continue
			} else if table[currentPoint[0]][currentPoint[1]-1] == "." && notSliceSearch(gonePoint, []int{currentPoint[0], currentPoint[1] - 1}) && notSliceSearch(prevPoint, []int{currentPoint[0], currentPoint[1] - 1}) {
				prevPoint = append(prevPoint, currentPoint)
				currentPoint = []int{currentPoint[0], currentPoint[1] - 1}
				continue
			} else {
				gonePoint = append(gonePoint, currentPoint)
				if notSliceCheck(currentPoint, start) {
					currentPoint = start
					continue
				}
				fmt.Println("NO")
				return
			}
		}
	}

}

func notSliceSearch(targetSleice [][]int, content []int) bool {
	if targetSleice == nil {
		return true
	}
	for _, r := range targetSleice {
		if r[0] == content[0] && r[1] == content[1] {
			return false
		}
	}
	return true
}

func notSliceCheck(prev []int, current []int) bool {
	if prev == nil {
		return true
	}
	if prev[0] == current[0] && prev[1] == current[1] {
		return false
	}
	return true
}
