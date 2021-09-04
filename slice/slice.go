package main

import "fmt"

//Tem um conceito parecido com arrays, porém não é necessario definir o seu tamanho
// Ponteiro: é o endereço de memória do array que será utilizado pelo slice para armazenar os dados.
// Tamanho: Quantidade de elementos que estão no slice.
// Capacidade: Quantidade de elementos que o array é capaz de armazenar. Essa capacidade muda dinamicamente de acordo com as alterações do slice.

func main() {
	slice := make([]int, 3, 5)  //slice tamanho 3 e capacidade 4
	slice2 := []int{}           //slice vazio
	slice3 := append(slice2, 2) //adicionando elemento no slice2

	fmt.Println("Tamanho slice 1:", len(slice))
	fmt.Println("Capacidade slice 1:", cap(slice))
	fmt.Println("Tamanho slice 2:", len(slice2))
	fmt.Println("Capacidade slice 2:", cap(slice2))
	fmt.Println("Tamanho slice 2 depois do append:", len(slice3))
	fmt.Println("Capacidade slice 2 depois do append:", cap(slice3))

}
