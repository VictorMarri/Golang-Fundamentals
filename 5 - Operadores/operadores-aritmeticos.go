package main

import "fmt"

func main() {
	soma := 1 + 2
	subtracao := 1 - 1
	divisao := 25 / 5
	multiplicacao := 12 * 3
	restoDaDivisao := 10 % 2

	fmt.Println(soma, subtracao, divisao, multiplicacao, restoDaDivisao)

	var numero1 int16 = 10
	var numero2 int16 = 20
	soma2 := numero1 + numero2 //Em GO nos so podemos fazer operacoes com numeros do mesmo tipo. Se eu tentar somar um int16 com um int32 o compilador nao vai deixar.
	fmt.Println(soma2)
}
