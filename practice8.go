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
	var count int

	for i := 0; i < height; i++ {
		sc.Scan()
		data := sc.Text()
		row = strings.Split(data, "")
		table = append(table, row)
	}

	for h := 0; h < height; h++ {
		for w := 0; w < width; w++ {
			if checkBuild(table, height, width, []int{h, w}) {
				count = count + 1
			}
		}
	}

	fmt.Println(count)
}

func checkBuild(table [][]string, height int, width int, building []int) bool {
	copyTable := make([][]string, height)
	for i := range table {
		copyTable[i] = make([]string, width)
		copy(copyTable[i], table[i])
	}
	copyTable[building[0]][building[1]] = "."
	var firstRake [][]int
	var point []int
	for h := 0; h < height; h++ {
		for w := 0; w < width; w++ {
			if copyTable[h][w] == "." {
				continue
			}
			if rakeSearch(firstRake, []int{h, w}) {
				continue
			}
			if copyTable[h][w] == "#" {
				if len(firstRake) > 0 {
					return false
				}
				point = []int{h, w}
				firstRake = append(firstRake, point)
				i := 0
				for {
					if i >= len(firstRake) {
						break
					}
					point = firstRake[i]
					if point[1]+1 < width && !rakeSearch(firstRake, []int{point[0], point[1] + 1}) && copyTable[point[0]][point[1]+1] == "#" {
						firstRake = append(firstRake, []int{point[0], point[1] + 1})
					}
					if point[0]+1 < height && !rakeSearch(firstRake, []int{point[0] + 1, point[1]}) && copyTable[point[0]+1][point[1]] == "#" {
						firstRake = append(firstRake, []int{point[0] + 1, point[1]})
					}
					if point[1]-1 >= 0 && !rakeSearch(firstRake, []int{point[0], point[1] - 1}) && copyTable[point[0]][point[1]-1] == "#" {
						firstRake = append(firstRake, []int{point[0], point[1] - 1})
					}
					if point[0]-1 >= 0 && !rakeSearch(firstRake, []int{point[0] - 1, point[1]}) && copyTable[point[0]-1][point[1]] == "#" {
						firstRake = append(firstRake, []int{point[0] - 1, point[1]})
					}
					i++
				}
			}
		}
	}
	return true
}

func rakeSearch(firstRake [][]int, point []int) bool {
	for _, r := range firstRake {
		if r[0] == point[0] && r[1] == point[1] {
			return true
		}
	}
	return false
}
