package main

import "fmt"

// ESTRUTURA BASICA DE UMA FUNCAO EM GO
// func {nomeFuncao}({parametro}{tipoDoParametro}){tipoDeRetornoDoMetodo}{
// }
func somar(numero1 int8, numero2 int8) int8 {
	return numero1 + numero2
}

func main() {
	valorSoma := somar(10, 20)
	fmt.Println(valorSoma)
}
