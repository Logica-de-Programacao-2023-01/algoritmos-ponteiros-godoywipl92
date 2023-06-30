package main

import "fmt"

type Livro struct {
	Titulo string
	Autor  string
}

func alterarTitulo(l *Livro) {
	if l.Autor == "Anônimo" {
		l.Titulo = "Desconhecido"
	}
}

func main() {
	livro := Livro{
		Titulo: "Livro1",
		Autor:  "Anônimo",
	}
	fmt.Println("Antes:", livro)

	alterarTitulo(&livro)
	fmt.Println("Depois:", livro)
}
