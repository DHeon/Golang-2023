package main

import "fmt"

func main() {
	// var primes [3]int
	// primes[0] = 2
	// primes[1] = 3
	// primes[2] = 5
	// fmt.Println(primes, primes[1])

	// var primes [3]int = [3]int{2, 3, 5}
	// fmt.Println(primes, primes[1])

	primes := [3]int{1, 2, 3}
	fmt.Println(primes, primes[1])

	test := [5]bool{true, true, true} //초기화 하지 않으면 해당 타입의 제로값으로 초기화됨.
	fmt.Println(test[3])

	for i := 0; i < 3; i++ {
		fmt.Println(primes[i])
	}

	i := 0
	for i < len(primes) {
		fmt.Println(primes[i])
		i++
	}
	fmt.Println("--------------------")
	for idx, prime := range primes {
		fmt.Println(idx, prime)
	}

	for prime := range primes { // 인덱스 값이 나옴.
		fmt.Println(prime)
	}
	//파이선 range 생각하면 안댐
	for _, prime := range primes { //인덱스  뺴고 값만.
		fmt.Println(prime)
	}

}
