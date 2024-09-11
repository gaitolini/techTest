package main

import "fmt"

func main() {
	// Faturamento mensal por estado
	faturamentoEstados := map[string]float64{
		"SP":     67836.43,
		"RJ":     36678.66,
		"MG":     29229.88,
		"ES":     27165.48,
		"Outros": 109849.53,
	}

	// Calcular o faturamento total
	var total float64
	for _, valor := range faturamentoEstados {
		total += valor
	}

	// Exibir o percentual de representação de cada estado
	fmt.Println("Percentual de representação por estado:")
	for estado, valor := range faturamentoEstados {
		percentual := (valor / total) * 100
		fmt.Printf("%s: %.2f%%\n", estado, percentual)
	}
}
