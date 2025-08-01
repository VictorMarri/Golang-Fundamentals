package main

import "fmt"

func main() {

	//Maps sao rigidos. Uma vez que voce declarou o tipo de dado que vai conter nos maps, vc so pode trabalhar com esse tipo e nenhum outro diferente
	//         map[tipoDasChaves]tipoDosValores{}
	usuario := map[string]string{
		"nome:":     "Victor",
		"sobrenome": "Marri",
	}

	fmt.Println(usuario["nome"])
}
