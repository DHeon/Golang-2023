package main

import "fmt"

func main() {
	var a int = 5
	double(&a)
	fmt.Printf("%d\n", a)
}

func double(number *int) {
	*number = *number * 2
}
