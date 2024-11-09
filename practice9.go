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

	for i := 0; i < height; i++ {
		sc.Scan()
		data := sc.Text()
		row = strings.Split(data, " ")
		table = append(table, row)
	}

	var pointList []int
	var fall struct {
		place int
		point int
	}
	var fallList []struct {
		place int
		point int
	}
	var nextFallList []struct {
		place int
		point int
	}

	for w := 0; w < width; w++ {
		fall.place = w
		firstpoint, _ := strconv.Atoi(table[0][w])
		fall.point = firstpoint
		fallList = append(fallList, fall)
		for h := 0; h < height; h++ {
			if h == height-1 {
				for i := 0; i < len(fallList); i++ {
					pointList = append(pointList, fallList[i].point)
				}
				fallList = nil
				break
			}
			for i := 0; i < len(fallList); i++ {
				if fallList[i].place-1 >= 0 {
					actvepoint, _ := strconv.Atoi(table[h+1][fallList[i].place-1])
					fall.point = fallList[i].point + actvepoint
					fall.place = fallList[i].place - 1
					nextFallList = append(nextFallList, fall)
				}
				actvepoint, _ := strconv.Atoi(table[h+1][fallList[i].place])
				fall.point = fallList[i].point + actvepoint
				fall.place = fallList[i].place
				nextFallList = append(nextFallList, fall)
				if fallList[i].place+1 < width {
					actvepoint, _ := strconv.Atoi(table[h+1][fallList[i].place+1])
					fall.point = fallList[i].point + actvepoint
					fall.place = fallList[i].place + 1
					nextFallList = append(nextFallList, fall)
				}
			}
			fallList = nextFallList
			nextFallList = []struct {
				place int
				point int
			}{}
		}
	}

	maxPoint := 0
	for i := 0; i < len(pointList); i++ {
		if maxPoint < pointList[i] {
			maxPoint = pointList[i]
		}
	}

	fmt.Println(maxPoint)
}
