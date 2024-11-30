// Stair段の階段をstep1段かstep2段とばしで進むことができる。
// その時に踏まない段はいくつあるかの問題
// （stair段を越えたると最後の段を必ず踏む）
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
	stair, _ := strconv.Atoi(sc.Text())
	sc.Scan()
	steps := strings.Split(sc.Text(), " ")
	step1, _ := strconv.Atoi(steps[0])
	step2, _ := strconv.Atoi(steps[1])

	var goneStair []int
	if step1 <= stair{
		goneStair = append(goneStair, step1)
	}
	if step2 <= stair{
		goneStair = append(goneStair, step2)
	}
	var gone, gone1, gone2 int
	stairMap := make(map[int]bool)
	stairMap[step1] = true
	stairMap[step2] = true
	
	for j := 0; j < len(goneStair); j++{
		gone = goneStair[j]
		gone1 = gone + step1
		gone2 = gone + step2
		if gone1 <= stair{
			if !stairMap[gone1]{
				stairMap[gone1] = true
				goneStair = append(goneStair, gone1)
			}
		}
		if gone2 <= stair{
			if !stairMap[gone2]{
				stairMap[gone2] = true
				goneStair = append(goneStair, gone2)
			}
		}
	}
	if !stairMap[stair]{
		goneStair = append(goneStair, stair)
	}
	fmt.Println(stair - len(goneStair))
}
