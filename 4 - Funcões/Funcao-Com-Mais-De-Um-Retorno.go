package main

import "fmt"

func calculosMatematicos(numero1, numero2 int8) (int8, int8) {
	soma := numero1 + numero2
	subtracao := numero1 - numero2
	return soma, subtracao
}

func main() {
	var soma, subtracao = calculosMatematicos(10, 20)
	fmt.Println(soma, subtracao)

	//Chamando funcoes com dois retornos.
	resultadoSoma, resultadoSubtracao := calculosMatematicos(10, 15)
	fmt.Println(resultadoSoma, resultadoSubtracao)

	//E quando eu nao preisar de um dos retornos?
	//Coloca um underline no retorno que voce nao precisa. Exemplo:
	resultadoSoma2, _ := calculosMatematicos(10, 15)
	fmt.Println(resultadoSoma2)

	_, resultadoSubtracao2 := calculosMatematicos(10, 15)
	fmt.Println(resultadoSubtracao2)
}
