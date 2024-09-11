
---

### 📂 **README para JavaScript**

```markdown
# 🚀 Calculando o percentual de faturamento por estado (JavaScript)

## 🛠️ O que é o PlayCode.io?
O **PlayCode.io** é uma ferramenta online onde você pode escrever, executar e testar código JavaScript diretamente no navegador, sem precisar instalar nada.

### 🔗 Acesse o PlayCode.io aqui: [PlayCode.io](https://playcode.io/javascript)

---

## 📋 Passo a Passo no PlayCode.io

### 1️⃣ **Acesse o link do PlayCode.io**
- Abra o navegador e vá até o **[PlayCode.io](https://playcode.io/javascript)**.

### 2️⃣ **Remova o código padrão**
- No editor que aparece na tela, apague o código padrão para que possamos inserir o nosso código.

### 3️⃣ **Cole o código JavaScript abaixo no editor**:

```javascript
// const faturamentoEstados = {
//     "AC": 15000.00,
//     "AL": 20000.00,
//     "AP": 12000.00,
//     "AM": 18000.00,
//     "BA": 45000.00,
//     "CE": 30000.00,
//     "DF": 35000.00,
//     "ES": 27165.48,
//     "GO": 40000.00,
//     "MA": 25000.00,
//     "MT": 22000.00,
//     "MS": 23000.00,
//     "MG": 29229.88,
//     "PA": 26000.00,
//     "PB": 21000.00,
//     "PR": 50000.00,
//     "PE": 38000.00,
//     "PI": 15000.00,
//     "RJ": 36678.66,
//     "RN": 19000.00,
//     "RS": 46000.00,
//     "RO": 13000.00,
//     "RR": 8000.00,
//     "SC": 40000.00,
//     "SP": 67836.43,
//     "SE": 14000.00,
//     "TO": 10000.00,
//     "Outros": 19849.53  // Caso inclua outros faturamentos de fora dos estados.
// };

const faturamentoEstados = {
    "SP": 67836.43,
    "RJ": 36678.66,
    "MG": 29229.88,
    "ES": 27165.48,
    "Outros": 19849.53
};

let total = 0;
for (let estado in faturamentoEstados) {
    total += faturamentoEstados[estado];
}

console.log("Percentual de representação por estado:");
for (let estado in faturamentoEstados) {
    let percentual = (faturamentoEstados[estado] / total) * 100;
    console.log(`${estado}: ${percentual.toFixed(2)}%`);
}
```
### 4️⃣ Execute o código 🏃‍♂️
- O PlayCode.io executa o código automaticamente. Se não executar, clique no botão Run (Executar) na parte superior direita da página.
### 5️⃣ Veja o resultado no painel de saída 🎉
- O programa exibirá o percentual de participação de cada estado no total de faturamento da distribuidora.
### 🔄 O que aprendemos?
- Como usar um objeto faturamentoEstados para armazenar os valores de faturamento por estado em JavaScript.
- Como calcular o percentual de representação de cada estado em relação ao valor total.
- Como formatar e exibir esses percentuais de maneira legível no console.
---
### 🚀 Executando o código localmente
Caso prefira executar o código no seu computador:

- Instale o Node.js.
- Crie um arquivo chamado index.js e cole o código acima.
- No terminal, navegue até o diretório onde o arquivo está salvo.
- Execute o programa com o comando:

~~~~ bash
- node index.js
~~~~

## 📲 Contato
Caso tenha dúvidas ou sugestões, entre em contato comigo:

 - Email: gaitolini@gmail.com
 - LinkedIn: [Anderson Gaitolini](https://www.linkedin.com/in/andersongaitolini/)
 - Whatsapp: [Me adicione no WhatsApp](https://wa.me/qr/CFND4RGOJHHUN1)
