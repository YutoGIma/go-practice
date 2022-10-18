package main

import "fmt"

func main() {
	fmt.Println("GO言語を練習しよう")
	
	num := 123 //var num int = 123の略
	str := "String" //var str string "String"の略
	
	fmt.Print("num=", num, "str=", str, "\n") //基本的な出力
	fmt.Println("num=", num, "str=", str) //自動的に空白と改行が入る
	fmt.Printf("num=%d str=%s\n", num, str) //printfの法則に従う
}