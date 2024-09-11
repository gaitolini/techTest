# 🚀 Calculando a sequência de Fibonacci e verificando se o número pertence à sequência (Go)

## 🛠️ O que é o Go Playground?
O **Go Playground** é uma ferramenta online onde você pode escrever, compilar e executar código Go diretamente no navegador, sem precisar instalar nada no seu computador. Muito útil para testar pequenas funções ou aprender Go!

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
    numero = 13 //Obs.: Como no playground não lê input, altere a variável no código 


    if fibonacci(numero) {
        fmt.Printf("O número %d pertence à sequência de Fibonacci.\n", numero)
    } else {
        fmt.Printf("O número %d não pertence à sequência de Fibonacci.\n", numero)
    }
}
```

## 📲 Contato
Caso tenha dúvidas ou sugestões, entre em contato comigo:

 - Email: gaitolini@gmail.com
 - LinkedIn: [Anderson Gaitolini](https://www.linkedin.com/in/andersongaitolini/)
 - Whatsapp: [Me adicione no WhatsApp](https://wa.me/qr/CFND4RGOJHHUN1)
