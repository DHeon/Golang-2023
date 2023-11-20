package main

import "fmt"

func main() {
	var a []string
	var b []bool //현재 공간이 없고 선언만 되있어서 nil상태임.
	// a = make([]string, 4, 5)

	b = append(b, true) //append하면 생김.
	fmt.Printf("%#v %#v\n", a, b)
	fmt.Println(a, len(a), cap(a))
	fmt.Println(b, len(b), cap(b))
}
