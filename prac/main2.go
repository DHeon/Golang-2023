package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	inputScore, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	inputScore = strings.TrimSpace(inputScore) //Trim 해줘야 ParseFloat 가능.
	is, err := strconv.ParseFloat(inputScore, 32)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(is > 3.5)

}
