package main

import "fmt"

func main() {
	// Exemplo de map simples
	usuario := map[string]string{
		"nome":      "Victor",
		"sobrenome": "Marri",
	}

	fmt.Println("Nome do usuário:", usuario["nome"])

	// Exemplo de map aninhado, simulando um registro mais complexo
	usuario2 := map[string]map[string]string{
		"pessoa": {
			"nome":      "Jane",
			"sobrenome": "Doe",
		},
		"curso": {
			"nome":   "Sistemas de Informação",
			"campus": "Barreiro",
		},
	}

	fmt.Println("Nome da pessoa:", usuario2["pessoa"]["nome"])
	fmt.Println("Sobrenome da pessoa:", usuario2["pessoa"]["sobrenome"])

	// Deletando uma chave do map
	fmt.Println("Antes de deletar 'pessoa':", usuario2)
	delete(usuario2, "pessoa")
	fmt.Println("Depois de deletar 'pessoa':", usuario2)

	// Adicionando uma nova chave ao map
	usuario2["cpf"] = map[string]string{
		"cpf": "12345678978",
	}
	fmt.Println("Após adicionar 'cpf':", usuario2)
}
