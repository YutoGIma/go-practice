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
	maxCount, _ := strconv.Atoi(sc.Text())
	sc.Scan()
	days := strings.Split(sc.Text(), " ")

	var last int
	var countList []int
	holidayCount := 0
	consecutiveCount := 0

	for i, _ := range days {
		last = i + 6
		for j := i; j <= last; j++ {
			if days[j] == "0" {
				holidayCount = holidayCount + 1
			}
		}
		if holidayCount >= 2 {
			consecutiveCount = consecutiveCount + 1
		} else {
			countList = append(countList, consecutiveCount)
			consecutiveCount = 0
		}
		holidayCount = 0
		if last == maxCount-1 {
			countList = append(countList, consecutiveCount)
			maxPoint := 0
			for i := 0; i < len(countList); i++ {
				if maxPoint < countList[i] {
					maxPoint = countList[i]
				}
			}
			if maxPoint != 0 {
				fmt.Println(maxPoint + 6)
			} else {
				fmt.Print(0)
			}
			return
		}
	}

}
