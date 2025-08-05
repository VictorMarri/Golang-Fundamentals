package main

import "fmt"

// Primeira forma de fazer um switch case
func diaDaSemana(numero int) string {
	switch numero {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda"
	case 3:
		return "Terca feira"
	case 4:
		return "Quarta feira"
	case 5:
		return "Quinta feira"
	case 6:
		return "Sexta feira"
	case 7:
		return "Sabado"
	default:
		return "Número inválido"
	}
}

// Outra forma de fazer um switch case
func diaDaSemana2(numero int) string {
	switch {
	case numero == 1:
		return "Domingo"
	case numero == 2:
		return "Segunda"
	case numero == 3:
		return "Terca feira"
	case numero == 4:
		return "Quarta feira"
	case numero == 5:
		return "Quinta feira"
	case numero == 6:
		return "Sexta feira"
	case numero == 7:
		return "Sabado"
	default:
		return "Número inválido"
	}
}

// Fallthrough clause
func diaDaSemana3(numero int) string {
	var diaDaSemana string

	switch {
	case numero == 1:
		diaDaSemana = "Domingo"
		fallthrough //Essa clausula joga meu codigo pra dentro da proxima condicao
	case numero == 2:
		diaDaSemana = "Segunda"
	case numero == 3:
		diaDaSemana = "Terca feira"
	case numero == 4:
		diaDaSemana = "Quarta feira"
	case numero == 5:
		diaDaSemana = "Quinta feira"
	case numero == 6:
		diaDaSemana = "Sexta feira"
	case numero == 7:
		diaDaSemana = "Sabado"
	default:
		diaDaSemana = "Número inválido"
	}

	return diaDaSemana
}

func main() {

	dia := diaDaSemana(3)
	fmt.Println(dia)

	fmt.Println("---------------")
	dia2 := diaDaSemana3(1)
	fmt.Println(dia2)
}
