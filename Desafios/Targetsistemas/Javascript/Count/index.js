function fibonacci(n) {
  let a = 0,
    b = 1;
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
