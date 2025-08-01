package main

import (
	"fmt"
	"reflect"
)

func main() {
	//Arrays

	var array1 [5]string
	fmt.Println(array1)

	array1[0] = "Hello"
	array1[1] = "World"
	fmt.Println(array1)

	//Array inferido
	array2 := [5]string{"Posicao0", "Posicao1", "Posicao2", "Posicao3", "Posicao4"}
	fmt.Println(array2)
	//array2[5] = "Posicao5"  -> o GO nao deixa voce adicionar itens a mais que o tamanho total do array

	//Uma forma mais `Flexivel` de popular arrays
	//Colocamos esses 3 pontos pra sinalizar q ele vai ser flexivel
	//Isso fixa o tamanho do array de acordo com a quantiade de itens que colocarmos entre as chaves
	//Nesse caso eu coloquei 5 itens, entao automaticamente o go infere que o tamanho vai ser 5
	array3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(array3)

	//FIM ARRAYS

	//SLICE

	//Estrutra feita baseada no array, porem mais flexivel. Principalmente com a questao de tamanho

	slice := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	fmt.Println(slice)

	fmt.Println(reflect.TypeOf(slice))  //Reflect devolve o tipo da variavel
	fmt.Println(reflect.TypeOf(array3)) //Devolve o tipo da variavel

	//Adicionando valores num slice
	slice = append(slice, 110) //Adiciona um item no slice, e retorna um slice novo, com as infos q ja existiam no slice + o item novo
	fmt.Println(slice)

	//Slice, como o proprio nome diz, é uma fatia. Portanto, podemos pegar uma FATIA do array. O Slice nada mais e que um ponteiro pra array
	slice2 := array2[1:3] //Pega os valores da posicao 1 a 2 do array2, e quando chegar no terceiro (3) ele para.
	fmt.Println(slice2)

	array2[1] = "Posicao alterada"
	fmt.Println(slice2)

	//FIM SLICE

	//Arrays internos
	slice3 := make([]float32, 10, 11) //Tipo, Tamanho total do array, capacidade maxima do array
	fmt.Println(slice3)
	fmt.Println("tamanho array:", len(slice3))    //Printando o tamanho do array
	fmt.Println("capacidade array:", cap(slice3)) //Printando a capacidade total do array

	//Quando nos criamos mais itens alem do `tamanho maximo` o go vai la, adiciona o item e dobra o tamanho do array
	slice3 = append(slice3, 5)
	slice3 = append(slice3, 6)
	fmt.Println(slice3)
	fmt.Println("tamanho array:", len(slice3))
	fmt.Println("capacidade array:", cap(slice3))

	slice4 := make([]float32, 5) //Slice de 5 posicoes
	fmt.Println(slice4)
	fmt.Println("tamanho slice4:", len(slice4))
	fmt.Println("capacidade slice4:", cap(slice4))

	slice4 = append(slice4, 10)
	fmt.Println("tamanho slice4:", len(slice4))
	fmt.Println("capacidade slice4:", cap(slice4))
	//FIM ARRAYS INTERNOS

}
