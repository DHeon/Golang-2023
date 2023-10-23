package main

import (
	"fmt"
	"log"
	"os"
)

func isPrime(n int) (bool, error) {
	if n < 2 {
		return false, fmt.Errorf("%d는 소수가 아닙니다", n)
	}
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false, nil
		}
	}
	return true, nil
}

func prime(n int) {
	p, err := isPrime(n)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	if p {
		fmt.Printf("%d는 소수입니다\n", n)
	} else {
		fmt.Printf("%d는 소수가 아닙니다\n", n)
	}
}

func primeRange(a int, b int) {
	if a > b {
		temp := a
		a = b
		b = temp
	}

	for i := a; i <= b; i++ {
		p, err := isPrime(i)
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}
		if p {
			fmt.Print(i, " ")
		}
	}
}
func main() {
	var menu int

	for true {
		fmt.Print("1) 소수판정 2) 구간 소수판정 : ")
		_, err := fmt.Scanln(&menu)
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}
		switch menu {
		case 1:
			var in int
			fmt.Print("정수 한 개 입력 :")
			_, err = fmt.Scanln(&in)

			if err != nil {
				log.Fatal(err)
			}
			prime(in)
		case 2:
			var a int
			var b int
			fmt.Print("정수 두 개 입력 : ")
			_, err = fmt.Scanln(&a, &b)
			if err != nil {
				log.Fatal(err)
			}
			primeRange(a, b)
		default:
			fmt.Print("프로그램을 종료합니다")
			os.Exit(1)
		}
	}
}
