package main

import "fmt"

func main() {
	// var a int //declaration
	// a = 7     //
	// fmt.Println(a)
	// var b = 8
	// fmt.Println(b)

	// c := 9
	// fmt.Println(c)

	// d := 8.34 //타입 추론으로 변수를 만들다 보니 float이 아닌 무조건 double로 만듬.
	// fmt.Println(d, reflect.TypeOf(d))

	// var e float32 = 8.34
	// fmt.Println(e, reflect.TypeOf(e))
	// //fmt.Println('Z', '2', '\n', '김', '인', '하') //rune 리터럴 int32
	// //fmt.Println(math.Floor(2.17), math.Ceil(2.17), math.Sqrt(16))
	// fmt.Println(reflect.TypeOf('Z'), reflect.TypeOf('2'), reflect.TypeOf("Hi"), reflect.TypeOf(4.99), reflect.TypeOf(false))

	//fmt.Println(strings.Title("open source \tprogramming \n\"go\"!"))
	//fmt.Println("OpensourceProgramming", "GO")

	//var c string //변수명은 영어대소문자로 시작해야함.  대문자로 시작하는 것은 다른 패키지에서 액세스 가능함.
	//그래서 fmt.Println이 대문자 P로 시작하는겅

	//package 명어 main이면 실행 파일로써 받아들이는듯.
	// var d int
	// var e bool
	// var f float64  //0이나 빈칸으로써 나옴
	// var G = 99

	// koreanzombie := "정찬성"
	// koreanZombie := "정찬성" //camel 표기

	// fmt.Println(koreanZombie)

	// a := 7
	// var b float32 = 8.34
	// var c float64 = 8.64
	// fmt.Println(a * int(b))
	// fmt.Println(c > b)  // 64 32도 안됨

	var price int = 100
	fmt.Println("Price is ", price, "dollars.")

	var taxRate float64 = 0.08
	var tax float64 = float64(price) * taxRate
	fmt.Println("Tax is", tax, "dollars.")

	var total float64 = float64(price) + tax
	fmt.Println("Total cost is", total, "dollars.")

	var availableFunds int = 120
	fmt.Println(availableFunds, "dollars available.")
	fmt.Println("Withnin budget?", float64(availableFunds) > total)
}
