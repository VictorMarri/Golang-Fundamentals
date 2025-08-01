package main

import "fmt"

// Structs (mesma coisa de classes em c#)
type usuario struct {
	nome     string
	idade    int8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     int8
}

//FIM Criação de structs

func main() {
	var u usuario
	fmt.Println(u) //Se nos nao passamos valores, ele vai inferir que todos os campos sao 0

	//Populando
	u.nome = "James"
	u.idade = 42
	fmt.Println(u)

	//Populando propriedades compostas
	enderecoExemplo := endereco{"Rua dos bobos", 0}

	//Outra forma de declarar, usando inferencia de tipo
	usuario2 := usuario{"David", 20, enderecoExemplo}

	fmt.Println(usuario2)

	//Criando usuario sem ter todos os dados de cara.
	usuario3 := usuario{
		nome: "James",
	}

	//A gente usa esse alias antes de declarar a variavel. Dessa forma a gente consegue declarar sem ter todos os valores

	fmt.Println(usuario3)
}
