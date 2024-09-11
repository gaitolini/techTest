package main

import "fmt"

func main() {
	indice := 13
	soma := 0
	k := 0

	for k < indice {
		k = k + 1
		soma = soma + k
	}

	// Exibe uma mensagem no terminal
	fmt.Println("O valor final da soma dos números de 1 a", indice, "é:", soma)
}
