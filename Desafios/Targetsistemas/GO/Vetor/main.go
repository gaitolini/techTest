package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	// JSON com dados de faturamento diário
	// "valor": 0.0 - Dias 3 e 11 Sem faturamento
	//"valor": 0.0 - Dias 6 e 7 Final de semana (sem faturamento)

	dadosJson := `[
        {"dia": 1, "valor": 2200.0},
        {"dia": 2, "valor": 1800.5},
        {"dia": 3, "valor": 0.0},   
        {"dia": 4, "valor": 2100.0},
        {"dia": 5, "valor": 2500.0},
        {"dia": 6, "valor": 0.0}, 
        {"dia": 7, "valor": 990.0},    
        {"dia": 8, "valor": 3200.0},
        {"dia": 9, "valor": 2300.0},
        {"dia": 10, "valor": 2800.0},
        {"dia": 11, "valor": 0.0},
        {"dia": 12, "valor": 2300.0},
        {"dia": 13, "valor": 3100.0}
    ]`

	finaisDeSemanaOuFeriados := map[int]bool{
		6: true, 7: true, // Dia 6 e 7 são finais de semana
	}

	type Faturamento struct {
		Dia   int     `json:"dia"`
		Valor float64 `json:"valor"`
	}

	var faturamento []Faturamento
	err := json.Unmarshal([]byte(dadosJson), &faturamento)
	if err != nil {
		log.Fatalf("Erro ao fazer o parse do JSON: %v", err)
	}

	if len(faturamento) == 0 {
		log.Fatalf("O slice de faturamento está vazio!")
	}

	var total, media float64
	var menor, maior float64
	diasAcimaDaMedia := 0
	diasComFaturamento := 0
	primeiroFaturamento := false

	for _, dado := range faturamento {
		if dado.Valor > 0 && !finaisDeSemanaOuFeriados[dado.Dia] {
			if !primeiroFaturamento {
				menor = dado.Valor
				maior = dado.Valor
				primeiroFaturamento = true
			}
			diasComFaturamento++
			total += dado.Valor
			if dado.Valor < menor {
				menor = dado.Valor
			}
			if dado.Valor > maior {
				maior = dado.Valor
			}
		}
	}

	if diasComFaturamento > 0 {
		media = total / float64(diasComFaturamento)
	}

	for _, dado := range faturamento {
		if dado.Valor > media {
			diasAcimaDaMedia++
		}
	}

	if primeiroFaturamento {
		fmt.Printf("Menor valor de faturamento: R$ %.2f\n", menor)
		fmt.Printf("Maior valor de faturamento: R$ %.2f\n", maior)
		fmt.Printf("Número de dias com faturamento acima da média mensal: %d\n", diasAcimaDaMedia)
	} else {
		fmt.Println("Nenhum faturamento registrado.")
	}
}
