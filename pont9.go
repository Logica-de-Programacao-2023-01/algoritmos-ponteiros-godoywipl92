package main

import "fmt"

type Livro struct {
	Titulo string
	Autor  string
	Preco  float64
}

func aplicarDesconto(l *Livro) {
	desconto := 0.10
	l.Preco -= l.Preco * desconto
}

func main() {
	livro := &Livro{
		Titulo: "O Grande Gatsby",
		Autor:  "F. Scott Fitzgerald",
		Preco:  39.99,
	}

	fmt.Println("Preço antes do desconto:", livro.Preco)
	aplicarDesconto(livro)
	fmt.Println("Preço depois do desconto:", livro.Preco)
}
