package main

import "fmt"

func main() {
	balls := make(map[string]int) //map[string]int{} 가 나옴
	// var balls map[string]int //map[string]int(nil)이 나옴 리터럴이나 make로 초기화가 안 됐기 떄문에 nil임.
	fmt.Printf("%#v\n", balls)

	balls["성기훈"] = 20
	fmt.Println(balls["성기훈"])
	fmt.Println(balls["오일남"]) //값 타입이 int이기 떄문에 int의 제로값 0이 뜸

}
