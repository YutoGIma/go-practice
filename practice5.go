package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()

	alphabetList := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
		"e": 5,
		"f": 6,
		"g": 7,
		"h": 8,
		"i": 9,
		"j": 10,
		"k": 11,
		"l": 12,
		"m": 13,
		"n": 14,
		"o": 15,
		"p": 16,
		"q": 17,
		"r": 18,
		"s": 19,
		"t": 20,
		"u": 21,
		"v": 22,
		"w": 23,
		"x": 24,
		"y": 25,
		"z": 26,
	}

	twoNames := strings.Split(sc.Text(), " ")
	firstName := twoNames[0]
	secondName := twoNames[1]

	firstNameList := strings.Split(firstName, "")
	secondNameList := strings.Split(secondName, "")

	firstNumberList := changeNumber(firstNameList, alphabetList)
	secondNumberList := changeNumber(secondNameList, alphabetList)

	number1 := addNumbers(append(firstNumberList, secondNumberList...))
	number2 := addNumbers(append(secondNumberList, firstNumberList...))

	if number1 > number2 {
		fmt.Println(number1)
	} else {
		fmt.Println(number2)
	}
}

func changeNumber(nameList []string, alphabetList map[string]int) []int {
	var numberList []int
	for i := 0; i < len(nameList); i++ {
		numberList = append(numberList, alphabetList[nameList[i]])
	}
	return numberList
}

func addNumbers(endNumberList []int) int {
	var appendList []int
	for {
		if len(endNumberList) == 1 {
			break
		}
		appendList = endNumberList
		endNumberList = nil
		for n := 0; n < (len(appendList) - 1); n++ {
			if (appendList[n] + appendList[(n+1)]) > 101 {
				endNumberList = append(endNumberList, (appendList[n] + appendList[(n+1)] - 101))
			} else {
				endNumberList = append(endNumberList, (appendList[n] + appendList[(n+1)]))
			}
		}
	}
	return endNumberList[0]
}
