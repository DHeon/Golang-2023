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
	fmt.Print("Input Score: ")
	reader := bufio.NewReader(os.Stdin)
	inputScoreStrings, err := reader.ReadString('\n') // option 2
	if err != nil {
		log.Fatal(err)
	}
	inputScoreStrings = strings.TrimSpace(inputScoreStrings)     //remove space
	inputScore, err := strconv.ParseFloat(inputScoreStrings, 32) // string to 32bit float
	var grade string
	if inputScore >= 90 {
		grade = "A grade!"
	} else {
		grade = "under A grade.."
	}
	fmt.Println("You got", grade)
}
