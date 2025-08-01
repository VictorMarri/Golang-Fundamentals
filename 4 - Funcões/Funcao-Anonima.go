package main

import "fmt"

//Funcao anonima é aquela funcao sem nome, atribuida diretamente a uma variavel ou usada como argumento

func main() {
	var f = func(txt string) {
		fmt.Println(txt)
	}

	f("Texto aqui")
}
