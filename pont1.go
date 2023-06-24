package main

import (
	"fmt"
)

func somaNaturais(p *int, n int) {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	*p = sum
}

func main() {
	numero := 0
	n := 5
	somaNaturais(&numero, n)
	fmt.Println("A soma dos", n, "primeiros números naturais é:", numero)
}
