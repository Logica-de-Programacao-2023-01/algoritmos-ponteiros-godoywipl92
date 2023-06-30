package main

import "fmt"

func modificarValor(p *int) {
	*p = 42
}
func main() {
	
	valor := new(int)
	fmt.Println("Valor antes:", *valor)

	modificarValor(valor)
	fmt.Println("Valor depois:", *valor)
}
func free(p *int) {
	
	*p = 0
}
