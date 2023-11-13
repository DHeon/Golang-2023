package main

import "fmt"

func main() {
	// s := make([]int, 5) // make로 구현. 변수 선언돠 동시에 메모리 할당, 제로값으로 초기화

	s := []int{0, 0, 0, 0, 0} //변수 선언과 동시에 메모리 할당 , 원소들 초기화.

	s[4] = 100
	s[0] = 7
	s[1] = 91
	for _, value := range s {
		fmt.Println(value)
	}

}
