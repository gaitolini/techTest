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
  SP: 67836.43,
  RJ: 36678.66,
  MG: 29229.88,
  ES: 27165.48,
  Outros: 19849.53,
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
