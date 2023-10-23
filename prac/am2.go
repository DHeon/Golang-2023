package main

import (
	"fmt"
	"os"
)

func main() {
	var menu int
	_, err := fmt.Scanln(&menu)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
}
