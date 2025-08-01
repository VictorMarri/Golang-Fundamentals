package auxiliar

import "fmt"

func escrever2() {
	fmt.Println("Escrevendo pela funcao escrever2")
}

//Como ela esta escrita em minusculo, eu so vou poder usar essa funcao dentro dos pacotes em que esta inserida
//Ou seja, tudo que estra dentro do pacote `auxiliar` pode usar 'escrever2'
//Mas, tudo que ta de fora, ou seja, por exemplo, o main.go não pode usar esse escrever2

//Se eu usar escrever2 dentro de auxiliar.go, e rodar esse auxiliar la no main, vai aparecer o println
//Se eu tentar usar escrever2 dentro do main.go ele sequer vai deixar
