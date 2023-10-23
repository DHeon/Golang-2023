package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("HI")
	// var a int
	// a = 7
	// fmt.Println(a)
	// var b = 7
	// fmt.Println(b)
	// c := 9
	// fmt.Println(c)

	// fmt.Println(reflect.TypeOf(c))
	fmt.Println('Z', '2', '\n', '김', '인', '하')
	fmt.Println(reflect.TypeOf(3.99)) // 얘도 어떻게 보면 타입 추론이라 64로 나오게 됨.
	fmt.Println("HI", "EVRO")

	a := 7
	var b float32 = 8.34
	var c float64 = 8.64
	fmt.Println(a * int(b))     //타입이 맞아야 연산가능.
	fmt.Println(c > float64(b)) // 64 32도 안됨 float64와 float32도 엄연히 타입이 달라서 비교가 안됨.

	fmt.Println(float32(a) > 3.9) // a가 7일때 a > 3.9 안됨. a가 7.1일때 a > 3 됨.???

	//float 변수 와 리터널 정수 비교는 괜찮지만
	//int 변수와 리터널 실수 비교는 안됨.
}
