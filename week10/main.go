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

//go install {모듈 이름}을 하면 exe가 만들어짐.

//지금은 go install main              사용자/fare1/go/bin 에 exe가 ㅁ나들어져있으.ㅁ
