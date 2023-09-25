package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Guess Number Game")
	answer := rand.Intn(100) + 1
	fmt.Println(answer)

	for i := 0; i < 10; i++ {
		fmt.Println("Left chance is ", 10-i)
		fmt.Print("Input Number: ")
		inputGuessString, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		inputGuessString = strings.TrimSpace(inputGuessString)
		inputGuess, err := strconv.Atoi(inputGuessString)
		if err != nil {
			log.Fatal(err)
		}
		if inputGuess == answer {
			fmt.Println("Congratulation")
			break
		} else if inputGuess < answer {
			fmt.Println("Your guess number is lower than answer!")
		} else if inputGuess > answer {
			fmt.Println("Your guess number is higher than answer!")
		}
	}

}
