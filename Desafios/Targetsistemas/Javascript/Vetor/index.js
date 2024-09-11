const dadosJson = [
  { dia: 1, valor: 2200.0 },
  { dia: 2, valor: 1800.5 },
  { dia: 3, valor: 0.0 },
  { dia: 4, valor: 2100.0 },
  { dia: 5, valor: 2500.0 },
  { dia: 6, valor: 0.0 }, // Final de semana (sem faturamento)
  { dia: 7, valor: 10.0 }, // Final de semana (sem faturamento)
  { dia: 8, valor: 3200.0 },
  { dia: 9, valor: 1500.0 },
  { dia: 10, valor: 2800.0 },
  { dia: 11, valor: 0.0 }, // Sem faturamento
  { dia: 12, valor: 2300.0 },
  { dia: 13, valor: 3100.0 },
];

const finaisDeSemanaOuFeriados = [6, 7];

let total = 0;
let menor = null;
let maior = null;
let diasComFaturamento = 0;
let diasAcimaDaMedia = 0;

dadosJson.forEach((dado) => {
  if (dado.valor > 0 && !finaisDeSemanaOuFeriados.includes(dado.dia)) {
    if (menor === null || dado.valor < menor) {
      menor = dado.valor;
    }
    if (maior === null || dado.valor > maior) {
      maior = dado.valor;
    }

    total += dado.valor;
    diasComFaturamento++;
  }
});

let media = diasComFaturamento > 0 ? total / diasComFaturamento : 0;

dadosJson.forEach((dado) => {
  if (dado.valor > media) {
    diasAcimaDaMedia++;
  }
});

console.log(
  "Menor valor de faturamento: R$",
  menor.toLocaleString("pt-BR", { style: "currency", currency: "BRL" })
);
console.log(
  "Maior valor de faturamento: R$",
  maior.toLocaleString("pt-BR", { style: "currency", currency: "BRL" })
);
console.log(
  "Número de dias com faturamento acima da média mensal:",
  diasAcimaDaMedia
);
