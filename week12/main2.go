package main

import "fmt"

func main() {
	var s []int
	s = make([]int, 5) // 알아서 제로값이 들어감. int니까 0이 들어감.

	for _, value := range s {
		fmt.Println(value)
	}

}
