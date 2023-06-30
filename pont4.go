package main

import (
	"fmt"
)

func somaUltimosDigitos(p *int) {
	num := *p
	digito1 := num % 10
	num /= 10
	digito2 := num % 10
	*p = digito1 + digito2
}

func main() {
	numero := 1234
	somaUltimosDigitos(&numero)
	fmt.Println("Soma dos últimos dígitos:", numero)
}
