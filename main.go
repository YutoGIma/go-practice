package main

import (
	"fmt"
	"strconv" //文字列型⇔数値型
)

type Person struct {
	name string
	age int
}

type Team struct {
	name string
	size int
}

type Printable interface {
	GetData() (string, int)
}

func add(x int, y int) int {
	return x + y
}

func addMinus(x int, y int) (int, int) {
	return x + y, x - y
}

func allSum(x int, y ... int) int {
	sum := x
	for _, num := range y {
		sum = sum + num
		fmt.Print(sum, " ")
	}
	fmt.Print("\n")
	return sum
}

func (p *Person) SetPerson(name string, age int){
	p.name = name
	p.age = age
}

func (p Person) GetData() (string, int){
	return p.name, p.age
}

func (t *Team) SetTeam(name string, size int){
	t.name = name
	t.size = size
}

func (t Team) GetData() (string, int){
	return t.name, t.size
}

func PrintOut(p Printable) {
	fmt.Println(p.GetData())
}

func strToInt(a interface {}){
	fmt.Printf("%d\n", a.(int))
}

func main() {
	fmt.Println("GO言語を練習しよう")

	num := 123      //var num int = 123の略
	str := "String" //var str string "String"の略

	fmt.Print("num=", num, "str=", str, "\n") //基本的な出力
	fmt.Println("num=", num, "str=", str)     //自動的に空白と改行が入る
	fmt.Printf("num=%d str=%s\n", num, str)   //printfの法則に従う

	// 文字列型⇔数値型
	var strNum string
	strNum = strconv.Itoa(num)
	fmt.Print("strNum=", strNum, "\n")

	// 配列の定義（配列は予め個数が決まっている）
	array1 := [3]string{}
	array1[0] = "one"
	array1[1] = "two"
	array1[2] = "three"
	fmt.Println("array1=[", array1[0], array1[1], array1[2], "]")
	a1 := [3]string{"Red", "Green", "Blue"} //初期化
	fmt.Println("a1=", a1)

	// スライスの定義（初期化時に配列の長さを指定しないもの）
	array2 := []string{}
	array2 = append(array2, "one")
	array2 = append(array2, "two")
	fmt.Println("array2=[", array2[0], array2[1], "]")

	// 連想配列（マップ）の定義
	map1 := map[int]string{
		1: "ichi",
		2: "ni",
	}
	map1[3] = "san" //追加
	fmt.Println("map1=", map1)

	// 繰り返しと条件分岐
	for i := 1; i < 13; i++ {
		if i%6 == 0 {
			fmt.Print(i, "は6の倍数です。\n")
		} else if i%3 == 0 {
			fmt.Print(i, "は3の倍数です。\n")
		} else if i%2 == 0 {
			fmt.Print(i, "は2の倍数です。\n")
		} else {
			fmt.Print(i, "は6の倍数でも、2の倍数でも、3の倍数でもありません。\n")
		}
	}

	for i := 0; i <= 10; i++ {
		fmt.Print(i, "\n")
		switch i {
		case 1:
			fmt.Print(1, "\n")
		case 2:
			fmt.Print(22, "\n")
		case 5:
			fmt.Print(55555, "\n")
		case 8:
			goto eight
		}
	}
	eight :
		fmt.Print("it is eight \n")

	fmt.Print(add(2, 3), "\n")

	// addMinusの複数の返り値の内、2番目のものだけをminusNumberに代入する
	_, minusNumber := addMinus(2, 3)
	fmt.Print(minusNumber, "\n")

	fmt.Print(allSum(1, 2, 3, 4, 5), "\n")

	var p1 Person
	p1.SetPerson("Gima", 23)
	name, age := p1.GetData()
	fmt.Printf("%s(%d)\n", name, age)
	
	// 下記のような利用も可能
	p2 := Person{name: "Riko", age: 23}
	name, age = p2.GetData()
	fmt.Printf("%s(%d)\n", name, age)

	strToInt(12)
	// strToInt("da")

	p3 := Person{name: "佐藤", age: 28}
	t1 := Team{name: "砂糖チーム", size: 12}
	PrintOut(p3)
	PrintOut(t1)
}
