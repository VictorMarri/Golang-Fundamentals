package main

import "fmt"

func main() {
	verdadeiro, falso := true, false

	fmt.Println(verdadeiro && falso) //Se todos forem TRUE
	fmt.Println(verdadeiro || falso) //Se algum deles for TRUE
	fmt.Println(!verdadeiro)         //Negacao. Inverter o valor de uma variavel
}
