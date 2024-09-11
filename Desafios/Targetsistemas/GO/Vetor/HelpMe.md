# 🚀 Calculando o menor, maior faturamento e dias acima da média (Go) usando JSON como fonte de dados

## 🛠️ O que é o Go Playground?
O **Go Playground** é uma ferramenta online onde você pode escrever, compilar e executar código Go diretamente no navegador, sem precisar instalar nada no seu computador.

### 🔗 Acesse o Go Playground aqui: [Go Playground](https://play.golang.org/)

---

## 📋 Passo a Passo no Go Playground

### 1️⃣ **Acesse o link do Go Playground**
- Abra o navegador e vá até o **[Go Playground](https://play.golang.org/)**.

### 2️⃣ **Remova o código padrão**
- No editor que aparece na tela, apague o código padrão para que possamos inserir o nosso código.

### 3️⃣ **Cole o código Go abaixo no editor**:

```go
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
```
### 4️⃣ Execute o código 🏃‍♂️
Clique no botão Run (Executar) no topo da página para rodar o código.
### 5️⃣ Veja o resultado no painel de saída 🎉
O resultado do programa será exibido abaixo do editor. O programa retornará o menor, o maior valor de faturamento e quantos dias tiveram faturamento acima da média mensal.
### 🔄 O que aprendemos?
- Como carregar dados de faturamento de um JSON em Go.
- Como calcular o menor, maior valor de faturamento e dias acima da média.
- Como ignorar dias sem faturamento, como finais de semana e feriados.

## 🚀 Executando o código localmente
Caso prefira executar o código no seu computador:

- Instale o Go.
- Crie um arquivo chamado faturamento.go e cole o código acima.
- No terminal, navegue até o diretório onde o arquivo está salvo.
- Execute o programa com o comando:

~~~~bash
go run faturamento.go
~~~~

## 📲 Contato
Caso tenha dúvidas ou sugestões, entre em contato comigo:

 - Email: gaitolini@gmail.com
 - LinkedIn: [Anderson Gaitolini](https://www.linkedin.com/in/andersongaitolini/)
 - Whatsapp: [Me adicione no WhatsApp](https://wa.me/qr/CFND4RGOJHHUN1)
