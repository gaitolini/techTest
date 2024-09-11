package main

import (
	"fmt"
)

func main() {
	var input string
	fmt.Print("Digite uma string para ser invertida: ")
	fmt.Scanln(&input)

	inverted := ""

	for i := len(input) - 1; i >= 0; i-- {
		inverted += string(input[i])
	}

	fmt.Println("String original:", input)
	fmt.Println("String invertida:", inverted)
}
