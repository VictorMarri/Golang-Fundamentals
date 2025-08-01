package main

import "fmt"

func main() {

	//O GO nao trabalha com operadores ternarios, como algumas linguagens, tipo o C#
	//Temos que usar o bom e velho IF/ELSE
	//O GO preza pela simpliciadade.

	var numero int32
	var texto string

	if numero > 5 {
		texto = "Maior que 5"
	} else {
		texto = "Menor que 5"
	}

	fmt.Println(texto)

}
