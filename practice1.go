// ChatGPTからの出題
// 1.変数 x に整数値 5 を代入し、それを出力してください。答えは何ですか？
// 2.文字列変数 message に "Hello, World!" を代入し、それを出力してください。答えは何ですか？
// 3.1 から 10 までの整数を for ループを使って順番に出力してください。
// 4.以下の整数スライス（リスト）を作成し、それを出力してください：[1, 2, 3, 4, 5]。
// 5.if 文を使って、整数変数 num が偶数か奇数かを判定して、それを出力してください。答えは何ですか？
// これらの問題に挑戦して、それぞれの答えを教えてください。

package main

import (
  "fmt"
)

func main(){
  
  x := 5
  fmt.Println(x)

  message := "Hello, World!"
  fmt.Println(message)

  for i := 1; i <= 10; i++ {
    fmt.Println(i)
  }

  list := []int{1, 2, 3, 4, 5}
  fmt.Println(list)

  num := 5

  if num % 2 == 0 {
    fmt.Println("偶数")
  } else {
    fmt.Println("奇数")
  }
}
