# 🚀 Invertendo uma string (Go)

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
    "fmt"
)

func main() {
    var input string
    // fmt.Print("Digite uma string para ser invertida: ")
    // fmt.Scanln(&input)
    input = "HelloWolrd"

    inverted := ""

    for i := len(input) - 1; i >= 0; i-- {
        inverted += string(input[i])
    }

    fmt.Println("String original:", input)
    fmt.Println("String invertida:", inverted)
}

```
### 4️⃣ Execute o código 🏃‍♂️
- Clique no botão Run (Executar) no topo da página para rodar o código.
### 5️⃣ Veja o resultado no painel de saída 🎉
- O programa exibirá a string original e a string invertida no console.
### 🔄 O que aprendemos?
- Como solicitar a entrada de dados do usuário no Go.
- Como inverter uma string manualmente, sem o uso de funções prontas, utilizando um loop simples que percorre a string de trás para frente..

## 🚀 Executando o código localmente
Caso prefira executar o código no seu computador:

- Instale o Go.
- Crie um arquivo chamado **main.go** e cole o código acima.
- No terminal, navegue até o diretório onde o arquivo está salvo.
- Execute o programa com o comando:
~~~~bash
go run main.go
~~~~

---

## 📲 Contato
Caso tenha dúvidas ou sugestões, entre em contato comigo:

 - Email: gaitolini@gmail.com
 - LinkedIn: [Anderson Gaitolini](https://www.linkedin.com/in/andersongaitolini/)
 - Whatsapp: [Me adicione no WhatsApp](https://wa.me/qr/CFND4RGOJHHUN1)
