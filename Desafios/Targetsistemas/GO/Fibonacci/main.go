package main

import "fmt"

func fibonacci(n int) bool {
	a, b := 0, 1
	for b <= n {
		if b == n {
			return true
		}
		a, b = b, a+b
	}
	return false
}

func main() {
	var numero int
	fmt.Print("Informe um número: ")
	fmt.Scan(&numero)

	if fibonacci(numero) {
		fmt.Printf("O número %d pertence à sequência de Fibonacci.\n", numero)
	} else {
		fmt.Printf("O número %d não pertence à sequência de Fibonacci.\n", numero)
	}
}
