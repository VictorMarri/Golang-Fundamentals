package main

import "fmt"

// Criando structs
type pessoa struct {
	nome      string
	sobrenome string
	idade     int8
	altura    uint8
}

type estudante struct {
	pessoa    //Passa somente o nome do outro struct
	curso     string
	faculdade string
}

//FIM Structs

//Quando a gente aplica essa forma de 'Herança' em go a gente simplesmente coloca o nome do outro struct e nao passa o tipo

func main() {
	fmt.Println("Heranca")

	pessoa1 := pessoa{"Victor", "Marri", 25, 172}
	fmt.Println(pessoa1)

	estudante1 := estudante{pessoa1, "sistemas", "puc"}
	fmt.Println(estudante1)
	fmt.Println(estudante1.pessoa)
	fmt.Println(estudante1.curso)
	fmt.Println(estudante1.faculdade)
	fmt.Println(estudante1.pessoa, estudante1.curso)

}
