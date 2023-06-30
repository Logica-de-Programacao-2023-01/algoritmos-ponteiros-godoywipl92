package main

import (
	"fmt"
	"math"
)

func mediaPi(p *float64) {
	*p = (*p + math.Pi) / 2
}

func main() {
	numero := 3.14
	mediaPi(&numero)
	fmt.Println("Média com Pi:", numero)
}
