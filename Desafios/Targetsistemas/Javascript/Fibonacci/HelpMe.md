
---

### 📂 **Helpme para JavaScript**:

# 🚀 Calculando a sequência de Fibonacci e verificando se o número pertence à sequência (JavaScript)

## 🛠️ O que é o PlayCode.io?
O **PlayCode.io** é uma ferramenta online onde você pode escrever, executar e testar código JavaScript diretamente no navegador, sem precisar instalar nada. Muito prático para testar pequenos trechos de código!

### 🔗 Acesse o PlayCode.io aqui: [PlayCode.io](https://playcode.io/javascript)

---

## 📋 Passo a Passo no PlayCode.io

### 1️⃣ **Acesse o link do PlayCode.io**
- Abra o navegador e vá até o **[PlayCode.io](https://playcode.io/javascript)**.

### 2️⃣ **Remova o código padrão**
- No editor que aparece na tela, apague o código padrão para que possamos inserir o nosso código.

### 3️⃣ **Cole o código JavaScript abaixo no editor**:
```javascript
function fibonacci(n) {
    let a = 0, b = 1;
    while (b <= n) {
        if (b === n) {
            return true;
        }
        [a, b] = [b, a + b];
    }
    return false;
}

let numero = prompt("Informe um número:");

if (fibonacci(parseInt(numero))) {
    console.log(`O número ${numero} pertence à sequência de Fibonacci.`);
} else {
    console.log(`O número ${numero} não pertence à sequência de Fibonacci.`);
}
```

## 📲 Contato
Caso tenha dúvidas ou sugestões, entre em contato comigo:

 - Email: gaitolini@gmail.com
 - LinkedIn: [Anderson Gaitolini](https://www.linkedin.com/in/andersongaitolini/)
 - Whatsapp: [Me adicione no WhatsApp](https://wa.me/qr/CFND4RGOJHHUN1)
