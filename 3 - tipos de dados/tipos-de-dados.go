package main

import (
	"errors"
	"fmt"
)

func main() {
	numero := 10000000000000000
	fmt.Println(numero)

	var numero2 uint32 = 1000
	fmt.Println(numero2)

	//alias

	//INT32
	var numero3 rune = 12355
	fmt.Println(numero3)

	//UINT8
	var numero4 byte = 123
	fmt.Println(numero4)

	//NUMEROS REAIS
	var numeroRea1 float32 = 123.45
	fmt.Println(numeroRea1)

	var numeroRea2 float64 = 1232342345.45
	fmt.Println(numeroRea2)

	numeroRea3 := 1232342345.45
	fmt.Println(numeroRea3)

	//STRINGS
	var str string = "Texto"
	fmt.Println(str)

	str2 := "Texto2"
	fmt.Println(str2)

	//DEVOLVE O VALOR DO CARACTERE NA TABELA ASCII
	char := 'R'
	fmt.Println(char)

	texto := 5
	fmt.Println(texto)

	//Booleano
	var booleano bool = true
	fmt.Println(booleano)

	//Error
	var erro error = errors.New("Erro Interno")
	fmt.Println(erro)
}
