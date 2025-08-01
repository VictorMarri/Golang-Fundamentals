package main

import "fmt"

func main() {
	//Existem duas maneiras de declarar variaveis em go. A forma explicita e implicita.

	//Explicita (Voce ja infere o tipo da variavel durante a criacão da mesma
	var variavel1 string = "Variavel 1 Explícita"
	fmt.Println(variavel1)

	//Implicita (O proprio go identifica qual o tipo atraves do valor atribuido a ela)
	variavel2 := "Variavel 2 Implicita"
	fmt.Println(variavel2)

	//Declarando varias variaveis de uma vez so
	var (
		variavel3 string = "Variavel 3"
		variavel4 string = "Variavel 4"
	)

	fmt.Println(variavel3, variavel4)

	//Ou da pra fazer assim

	variavel5, variavel6 := "Variavel 5", "Variavel 6"
	fmt.Println(variavel5, variavel6)

	//Declarando constantes
	const constante1 string = "constante 1"
	fmt.Println(constante1)
}
