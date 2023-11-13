package main

import "fmt"

func main() {
	a := make([]string, 4, 5) // type, length, capacity 용량은 미리 주소를 확보하는것??
	// 용량은 현재 슬라이스에 주소값의 변경없이 포함시킬 수 있는 요소의 개수를 나타낸다.
	//값을 계속 append해서 length가 맨첨에 정한 capacity를 넘으면 다른 주소에 저장함.

	a[0] = "a"
	a[1] = "b"
	a[2] = "c"
	as := a[0:2]
	as[1] = "Z"
	// c := append(a, "y")
	c := append(a, "y", "x")

	fmt.Println(a, len(a), cap(a))
	fmt.Println(c, len(c), cap(c))
}
