package main

import "fmt"

func isPrimo(num int) bool {
	if num < 2 {
		return false
	}

	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}

	return true
}

func preencherPrimos(primes *[]int, n int) {
	count := 0
	num := 2

	for count < n {
		if isPrimo(num) {
			*primes = append(*primes, num)
			count++
		}
		num++
	}
}

func main() {
	var primes []int
	tamanho := 10

	preencherPrimos(&primes, tamanho)

	fmt.Println("Números primos:", primes)
}
