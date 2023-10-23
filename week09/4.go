package main

import "fmt"

func main() {
	a := 10
	b := 20

	c := &a

	fmt.Printf("%T\n", c)
	fmt.Println(a, b)
	swap(&a, &b)
	fmt.Println(a, b)

}

func swap(n1 *int, n2 *int) {
	temp := *n1
	*n1 = *n2
	*n2 = temp
}
