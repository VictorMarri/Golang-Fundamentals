package main

import "fmt"

func main() {
	var numero int = 100000000 //Se eu nao declarar o tipo de int (8,32,64) ele vai inferir automaticamente de acordo com os bits da minha maquina
	fmt.Println(numero)

	var numero2 uint32 = 1000
	fmt.Println(numero2)

	//RUNE = INT32
	var numero3 rune = 12456
	fmt.Println(numero3)

	//BYTE = uint8
	var numero4 byte = 123
	fmt.Println(numero4)
}
