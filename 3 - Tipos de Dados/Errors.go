package main

import (
	"errors" //PACOTE ESPECIFICO PRA TRATAR ERROS
	"fmt"
)

func main() {
	var erro error = errors.New("Erro interno")
	fmt.Println(erro)
}
