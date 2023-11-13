package main

import "fmt"

func main() {
	s := []int{0, 0, 0, 0, 0}

	s[4] = 100
	s[0] = 7
	s[1] = 91
	for _, value := range s {
		fmt.Println(value)
	}
	copyS := s[1:4]
	for _, value := range copyS {
		fmt.Println(value)
	}

	test := [3]string{"inha", "go", "student"}
	// fmt.Println(test, len(test))

	testS := test[:2]
	testS2 := test[1:]

	testS2[0] = "python"
	// test[1] = "python" 원본 또는 슬라이스를 건드림.

	fmt.Println(test, len(test))
	fmt.Println(testS, len(testS))
	fmt.Println(testS2, len(testS2))
}
