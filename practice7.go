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
	size, _ := strconv.Atoi(sc.Text())

	var table [][]string
	var row []string

	for i := 0; i < size; i++ {
		sc.Scan()
		data := sc.Text()
		row = strings.Split(data, "")
		table = append(table, row)
	}

	maxSwipLength := checkswipe(table, size)

	fmt.Println(maxSwipLength)
}

func checkswipe(table [][]string, size int) int {
	var point int
	var comparePoint int
	var diff int
	count := 0
	maxSwipLength := 0
	for h := 0; h < size; h++ {
		for w := 0; w < size; w++ {
			point, _ = strconv.Atoi(table[h][w])
			for i := 1; i < size; i++ {
				if (h - i) > 0 {
					comparePoint, _ = strconv.Atoi(table[(h - i)][w])
					if i == 1 {
						if abs(point-comparePoint) == 1 {
							count = count + 2
							diff = point - comparePoint
						} else {
							break
						}
					} else {
						if (point - comparePoint) == diff {
							count = count + 1
						} else {
							break
						}
					}
					point = comparePoint
				} else {
					break
				}
			}
			if maxSwipLength < count {
				maxSwipLength = count
			}
			diff = 0
			count = 0
			for i := 1; i < size; i++ {
				if (h + i) < size {
					comparePoint, _ = strconv.Atoi(table[(h + i)][w])
					if i == 1 {
						if abs(point-comparePoint) == 1 {
							count = count + 2
							diff = point - comparePoint
						} else {
							break
						}
					} else {
						if (point - comparePoint) == diff {
							count = count + 1
						} else {
							break
						}
					}
					point = comparePoint
				} else {
					break
				}
			}
			if maxSwipLength < count {
				maxSwipLength = count
			}
			diff = 0
			count = 0
			for i := 1; i < size; i++ {
				if (w - i) > 0 {
					comparePoint, _ = strconv.Atoi(table[h][(w - i)])
					if i == 1 {
						if abs(point-comparePoint) == 1 {
							count = count + 2
							diff = point - comparePoint
						} else {
							break
						}
					} else {
						if (point - comparePoint) == diff {
							count = count + 1
						} else {
							break
						}
					}
					point = comparePoint
				} else {
					break
				}
			}
			if maxSwipLength < count {
				maxSwipLength = count
			}
			diff = 0
			count = 0
			for i := 1; i < size; i++ {
				if (w + i) < size {
					comparePoint, _ = strconv.Atoi(table[h][(w + i)])
					if i == 1 {
						if abs(point-comparePoint) == 1 {
							count = count + 2
							diff = point - comparePoint
						} else {
							break
						}
					} else {
						if (point - comparePoint) == diff {
							count = count + 1
						} else {
							break
						}
					}
					point = comparePoint
				} else {
					break
				}
			}
			if maxSwipLength < count {
				maxSwipLength = count
			}
			diff = 0
			count = 0
			for w := 0; w < size; w++ {
				point, _ = strconv.Atoi(table[h][w])
				for i := 1; i < size; i++ {
					if (h-i) > 0 && (w-i) > 0 {
						comparePoint, _ = strconv.Atoi(table[(h - i)][(w - i)])
						if i == 1 {
							if abs(point-comparePoint) == 1 {
								count = count + 2
								diff = point - comparePoint
							} else {
								break
							}
						} else {
							if (point - comparePoint) == diff {
								count = count + 1
							} else {
								break
							}
						}
					} else {
						break
					}
					point = comparePoint
				}
				if maxSwipLength < count {
					maxSwipLength = count
				}
				diff = 0
				count = 0
				for i := 1; i < size; i++ {
					if (h+i) < size && (w+i) < size {
						comparePoint, _ = strconv.Atoi(table[(h + i)][(w + i)])
						if i == 1 {
							if abs(point-comparePoint) == 1 {
								count = count + 2
								diff = point - comparePoint
							} else {
								break
							}
						} else {
							if (point - comparePoint) == diff {
								count = count + 1
							} else {
								break
							}
						}
					} else {
						break
					}
					point = comparePoint
				}
				if maxSwipLength < count {
					maxSwipLength = count
				}
				diff = 0
				count = 0
				for i := 1; i < size; i++ {
					if (w-i) > 0 && (h+i) < size {
						comparePoint, _ = strconv.Atoi(table[(h + i)][(w - i)])
						if i == 1 {
							if abs(point-comparePoint) == 1 {
								count = count + 2
								diff = point - comparePoint
							} else {
								break
							}
						} else {
							if (point - comparePoint) == diff {
								count = count + 1
							} else {
								break
							}
						}
					} else {
						break
					}
					point = comparePoint
				}
				if maxSwipLength < count {
					maxSwipLength = count
				}
				diff = 0
				count = 0
				for i := 1; i < size; i++ {
					if (w+i) < size && (h-i) > 0 {
						comparePoint, _ = strconv.Atoi(table[(h - i)][(w + i)])
						if i == 1 {
							if abs(point-comparePoint) == 1 {
								count = count + 2
								diff = point - comparePoint
							} else {
								break
							}
						} else {
							if (point - comparePoint) == diff {
								count = count + 1
							} else {
								break
							}
						}
					} else {
						break
					}
					point = comparePoint
				}
				if maxSwipLength < count {
					maxSwipLength = count
				}
				diff = 0
				count = 0
			}
		}
	}
	return maxSwipLength
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
