# 🚀 Calculando o percentual de faturamento por estado (Go)

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

import "fmt"

func main() {
    // Faturamento mensal por estado
    faturamentoEstados := map[string]float64{
        "SP":     67836.43,
        "RJ":     36678.66,
        "MG":     29229.88,
        "ES":     27165.48,
        "Outros": 19849.53,
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
```

### 4️⃣ Execute o código 🏃‍♂️
Clique no botão Run (Executar) no topo da página para rodar o código.
### 5️⃣ Veja o resultado no painel de saída 🎉
O programa exibirá o percentual de participação de cada estado no total de faturamento da distribuidora.
## 🔄 O que aprendemos?
- Como usar um map para armazenar os valores de faturamento por estado em Go.
- Como calcular o percentual de representação de cada estado em relação ao valor total.
- Como formatar e exibir esses percentuais de maneira legível.

## 🚀 Executando o código localmente
### Caso prefira executar o código no seu computador:
- Instale o Go.
- Crie um arquivo chamado main.go e cole o código acima.
- No terminal, navegue até o diretório onde o arquivo está salvo.

### Execute o programa com o comando:
~~~~bash
go run main.go
~~~~

## 📲 Contato
Caso tenha dúvidas ou sugestões, entre em contato comigo:

 - Email: gaitolini@gmail.com
 - LinkedIn: [Anderson Gaitolini](https://www.linkedin.com/in/andersongaitolini/)
 - Whatsapp: [Me adicione no WhatsApp](https://wa.me/qr/CFND4RGOJHHUN1)
