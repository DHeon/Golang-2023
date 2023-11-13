package main

import "fmt"

func main() {
	s := []int{0,0,0,0,0}

	s[0] = s[4] = 100
	s[0] = 7
	s[1] = 91
	for _, value := range s {
		fmt.Println(value)
	}
	copyS := s[1:4]
	for _, value := range copyS{
		fmt.Println(value)
	}

	test := [3]string{"inha", "go", "student"}
	fmt.Println(test, len(test))

	coptS := test[:2];
	fmt.Println(test, len(test))
	fmt.Println(testS, len(testS))
}
