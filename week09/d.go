package main

import (
	"errors"
	"fmt"
)

func main() {
	err := fmt.Errorf("Dd")
	fmt.Print(err)
	fmt.Print(err.Error())
	e := errors.New("dld")
	fmt.Println(e)
	fmt.Println(e.Error())

}
