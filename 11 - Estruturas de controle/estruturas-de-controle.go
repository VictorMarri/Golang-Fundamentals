package main

import "fmt"

func main() {
	fmt.Println("Estruturas de controle em Go")

	numero := -5

	// Estrutura mais básica de if/else
	if numero > 15 {
		fmt.Println("Maior que 15")
	} else {
		fmt.Println("Menor ou igual a 15")
	}

	// If com declaração curta (if init)
	// Use if init quando a variável só serve para aquele teste e não será usada depois.
	if outroNumero := numero; outroNumero > 0 {
		fmt.Println("Número é maior que 0")
	} else if outroNumero < -10 {
		fmt.Println("Número é menor que -10")
	} else {
		fmt.Println("Número está entre 0 e -10")
	}

	// Exemplo prático de if init
	//checando se uma chave existe em um map
	idades := map[string]int{
		"João":  30,
		"Maria": 25,
	}

	// Aqui, idade e existe só existem dentro desse if
	if idade, existe := idades["João"]; existe {
		fmt.Println("Idade de João:", idade)
	} else {
		fmt.Println("João não encontrado no mapa de idades")
	}

	// Outro exemplo com uma chave inexistente
	if idade, existe := idades["Pedro"]; existe {
		fmt.Println("Idade de Pedro:", idade)
	} else {
		fmt.Println("Pedro não encontrado no mapa de idades")
	}
}
