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
		row = strings.Split(data, "")
		table = append(table, row)
	}

	var count int

	for h := 0; h < height; h++ {
		if h == 0 || h == (height-1) {
			continue
		}
		for w := 0; w < width; w++ {
			if w == 0 || w == (width-1) {
				continue
			}
			if table[h][w] == "." && table[h][(w+1)] == "#" && table[h][(w-1)] == "#" {
				if table[(h - 1)][(w-1)] != "#" || table[(h - 1)][(w)] != "#" || table[(h - 1)][(w+1)] != "#" {
					continue
				}
				if table[(h + 1)][(w-1)] != "#" || table[(h + 1)][(w)] != "#" || table[(h + 1)][(w+1)] != "#" {
					continue
				}
				count = count + 1
			}
		}
	}

	fmt.Print(count)
}
