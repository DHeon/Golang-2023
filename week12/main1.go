package main

import "fmt"

//슬라이스랑 배열 다름..
//슬라이스 동적 배열
//[]int  [] 안에 암 것도 안 넣으면 슬라이스임.x`

//a1 := [5]int{1,2,3,4,5}
//a2 := a1[0:3]
//a3 := a1[2:5]
//이거는 슬라이싱이라고 함 현재 배열을 기반으로 슬라이싱해서 슬라이스 만드는겅.
//참고로 원본을 바꾸면 슬라이싱 된애들ㄷ 바뀌고 슬라이싱 애들을 바꿔도 원본이 바뀜.
func main() {
	numbers := make([]float64, 3)
	numbers[0] = 19.7
	numbers[2] = 25.2
	for i, number := range numbers {
		fmt.Println(i, number)
	}
	var letters = []string{"a", "b", "c"}
	for i, letter := range letters {
		fmt.Println(i, letter)
	}
}
