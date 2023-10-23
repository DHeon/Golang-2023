package main

import "fmt"

func main() {
	a := 10
	b := 20

	c := &a
	//fmt.Printf("%d %d %x\n", a, b, c) //x로 하면 16진수로 나오는듯.
	fmt.Printf("%d %d %d\n", a, b, *c)

}

func swap(n1 *int, n2 *int) {
	temp := *n1
	*n1 = *n2
	*n2 = temp
}
