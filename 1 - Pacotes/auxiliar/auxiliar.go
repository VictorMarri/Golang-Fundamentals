package auxiliar

import "fmt"

// Registra uma mensagem na tela.
func Escrever() {
	fmt.Println("Escrevendo do arquivo auxiliar")
	escrever2() //Aqui, diferentemente do arquivo main, a gente consegue chamar esse escrever dois, pois ele ta dentro do mesmo pacote
}

//Se a funcao comecar com letra minuscula (Ex: escrever) ela é visível apenas dentro do pacote q ela esta
//Se comeca com letra maiuscula, ela pode ser exportada para outros pacotes

//Por padrao no go, quando temos uma funcao que esta sendo exportada, a gente tem que colocar um comentario em cima dela falando o que ela ta fazendo
