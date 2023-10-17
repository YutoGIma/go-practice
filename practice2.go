// Fibonacci 数列の最初の10個の数値を計算し、それらを出力してください。Fibonacci 数列は、最初の2つの数が1で、それ以降の数は直前の2つの数の和である数列です。
// 最初の10個の Fibonacci 数は次のようになります: 1, 1, 2, 3, 5, 8, 13, 21, 34, 55。

package main

import (
	"fmt"
)

func main() {
	fibonacci := []int{1, 1}

	for i := 0; i < 8; i++ {
		num := fibonacci[i] + fibonacci[i+1]
		fibonacci = append(fibonacci, num)
	}

	fmt.Println(fibonacci)
}
