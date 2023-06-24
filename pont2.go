package main

import (
	"fmt"
)

func verificaParImpar(p *int) {
	if *p%2 == 0 {
		*p = 0 
	} else {
		*p = 1 
	}
}

func main() {
	numero := 7
	verificaParImpar(&numero)
	fmt.Println("O número após a verificação é:", numero)
}
