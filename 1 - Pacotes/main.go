package main

import (
	"awesomeProject/auxiliar"
	"fmt"
	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Escrevendo do arquivo main")
	auxiliar.Escrever() //Aqui nos conseguimos chamar pq ela esta com letra maiuscula, ou seja, é visível pra outros pacotes
	erro := checkmail.ValidateFormat("victorpujoni@gmail.com")
	fmt.Println(erro)
}
