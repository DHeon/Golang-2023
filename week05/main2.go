package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Print("Input Score: ")
	reader := bufio.NewReader(os.Stdin)
	inputScore, _ := reader.ReadString('\n') // option 1
	fmt.Println(inputScore)
}
