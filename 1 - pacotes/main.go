package main

import (
	"fmt"
	"modulo/auxiliar"

	"rsc.io/quote"
)

func main() {
	fmt.Println("Escrevendo com GO!")
	auxiliar.Escrever()
	fmt.Println(quote.Glass())
}
