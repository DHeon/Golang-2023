package main

import (
	"fmt"
	"main/src/greeting"
	"main/src/mymath"
)

func main() {
	fmt.Println(mymath.MyPower(2, 9))
	fmt.Println(mymath.MyPower(4, 3))
	fmt.Println(mymath.MyAbs(-99))
	fmt.Println(mymath.MyAbs(71))
	greeting.Hi()
	greeting.Hello()
	greeting.Hello()
}
