package main

import "fmt"

func main() {
	fmt.Println("Ponteiros")

	//Ponteiros nao guardam valor, e sim referencias de memoria.

	var variavel1 int = 10
	var variavel2 int = variavel1

	fmt.Println(variavel1, variavel2)

	variavel1++
	//Quando formos imprimir o valor, vamos observar que o valor da variavel1 foi pra 11 e a variavel2 ficou em 10
	//Isso acontece porque quando vc atribui um valor a uma variavel, ele é uma copia.
	fmt.Println(variavel1, variavel2)

	//Atribuindo valores usando ponteiros
	//Ponteiros sao referencia de memora. Quando criamos um ponteiro e jogamos a variavel nele
	//Nos vamos estar jogando apenas o endereco de memoria onde a vairiavel ta salva
	//Exemplo:
	var variavel3 int //Guarda o valor do inteiro
	var ponteiro *int //Guarda o endereco de memoria de um valor inteiro

	fmt.Println(variavel3, ponteiro)

	variavel3 = 100
	ponteiro = &variavel3 //Aqui eu n to jogando o valor 100 pro ponteiro, eu to passando pra ele o endereço de memoria onde esse valor 100 ta armazenado

	fmt.Println(variavel3, *ponteiro) //Agora, se eu quiser ver o valor que ta salvo no endereco de memroria, eu coloco um asterisco na frente dele. Esse processo se chama desreferenciação

	variavel3 = 150
	fmt.Println(variavel3, ponteiro)
	fmt.Println(variavel3, *ponteiro)

	//Quando voce quer usar um mesmo valor em varios lugares, voce usa ponteiros. Dessa forma voce nao usa copia de valores, e sim o valor na sua raiz
}
