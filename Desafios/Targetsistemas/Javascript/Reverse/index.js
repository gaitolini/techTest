let input = "Distribuidora";
// let input = prompt("Digite uma string para ser invertida:");
let inverted = "";

for (let i = input.length - 1; i >= 0; i--) {
  inverted += input[i];
}

console.log("String original:", input);
console.log("String invertida:", inverted);
