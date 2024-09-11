# Tech Test
🧠 Testes Lógicos de Programação para Novo Emprego 🚀

Bem-vindo ao repositório de Testes Lógicos de Programação! 
Este repositório foi criado para compartilhar minhas habilidades em lógica de programação, especialmente para entrevistas técnicas. 
Aqui, você encontrará desafios de lógica implementados em diversas linguagens, focando em Golang, JavaScript, Delphi, SQL e até em pseudocódigos. 

Estes exercícios são demonstrar aos recrutadore as minhas habilidades de raciocínio lógico e programação enquanto me prepara para uma oportunidae!

✨ Objetivos do Projeto
 - 💻 Demonstrar habilidades de resolução de problemas: cada desafio aborda problemas comuns em entrevistas técnicas.
 - 🧠 Demonstrar lógica de programação: todos os exercícios envolvem algoritmos e estruturação lógica.
 - 🚀 Preparação para entrevistas: Enviar questões que simulam desafios técnicos de empresas de tecnologia.

📁 Estrutura do Repositório

O repositório está organizado por linguagens de programação. Cada diretório contém desafios e suas respectivas soluções. Abaixo, um exemplo da estrutura:
~~~~bash
/desafios-logica-programacao/
│
├── Golang/
│   ├── Hello-world/
│   │   ├──code.pas
│   ├── Resistor-color
│   │   ├──code.pas
│   └── README.md
│
├── JavaScript/
│   ├── Hello-world/
│   │   ├──code.pas
│   ├── Resistor-color
│   │   ├──code.pas
│   └── README.md
│
├── Delphi/
│   ├── Hello-world/
│   │   ├──code.pas
│   ├── Resistor-color
│   │   ├──code.pas
│   └── README.md
│
├── SQL/
│   ├── desafio1.sql
│   ├── desafio2.sql
│   └── README.md
│
├── Pseudocódigos 
│   ├── desafio1.sql
│   ├── desafio2.sql
│   └── README.md
│
└── Desafios/
    ├── desafio1.md
    ├── desafio2.md
    └── desafio3.md

~~~~

📝 Como Usar
Clone o repositório:

~~~~bash
git clone https://github.com/gaitolini/techTest
~~~~
Navegue até a pasta da linguagem desejada:

~~~~bash
cd Golang
~~~~

👩🏻‍💻Escolha um desafio e execute o arquivo:

~~~~bash
go run main.go
~~~~

Cada desafio possui um README explicando o problema, a solução e como executá-lo.

🚩 Desafios Incluídos
1. Panagrama (Golang)
Um pangrama é uma frase que utiliza todas as letras do alfabeto pelo menos uma vez.
**Exemplo de execução: "Jovem exibe fotografia da fazenda queimada com queijo, vinho, waffles e suco de uva."**

~~~~go
package pangram

import "strings"

func IsPangram(input string) bool {

	pangram := "abcdefghijklmnopqrstuvwxyz"
	input = strings.ToLower(input)
	for _, v := range pangram {
		if !strings.Contains(input, string(v)) {
			return false
		}
	}
	return true

}

~~~~
2. Verificação de Números Primos (JavaScript)
Escreva uma função que verifique se um número é primo.

Exemplo de execução: 17 é primo?

~~~~javascript
function isPrime(num) {
  if (num <= 1) return false; 
  if (num === 2) return true;
  if (num % 2 === 0) return false;

  for (let i = 3; i <= Math.sqrt(num); i += 2) {
    if (num % i === 0) return false;
  }
  return true;
}
~~~~

3. Neste exercício, você crio um programa útil para que não precise memorizar os valores das faixas de resistores. (Delphi)
Essas cores são codificadas da seguinte maneira:
~~~~bash
- ⚫preto: 0
- 🟤marrom: 1
- 🔴vermelho: 2
- 🟠laranja: 3
- 🟡amarelo: 4
- 🟢verde: 5
- 🔵azul: 6
- 🟣violeta: 7
- 🩶cinza: 8
- ⚪branco: 9
~~~~

Exemplo de execução:

~~~~Pascal
unit uResistorColor;

interface

uses
  System.SysUtils,
  System.Generics.Collections;

type
  TResistor = class
  private
    class var FColors: TArray<string>;
  public
    class property colors: TArray<string> read FColors;
    class function colorCode(const aColorName: string = ''): Integer;
    class constructor Create;
  end;

implementation

{ TResistor }

class constructor TResistor.Create;
begin
  FColors := TArray<string>.Create('black', 'brown', 'red', 'orange',
    'yellow', 'green', 'blue', 'violet', 'grey', 'white');
end;

class function TResistor.colorCode(const aColorName: string): Integer;
var
  i: Integer;
begin
  Result := -1;

  for i := Low(FColors) to High(FColors) do
  begin
    if FColors[i] = aColorName then
    begin
      Result := i;
      Exit;
    end;
  end;

  if Result = -1 then
    raise Exception.CreateFmt('Cor "%s" não encontrada no código de cores.', [aColorName]);
end;

end.
~~~~

4. Consulta de Funcionários por Salário (SQL)
Crie uma query SQL que retorne os funcionários com salários acima de um valor especificado.

Exemplo de execução:

~~~~sql
SELECT * FROM funcionarios WHERE salario > 5000;
~~~~

💡 Como Contribuir
Adoraria contar com sua contribuição! Siga os passos abaixo para me enviar novos desafios ou melhorias:

Faça um fork do repositório.
Crie uma nova branch para suas alterações:
~~~bash
git checkout -b nova-feature
~~~~
Faça commit das suas modificações:

~~~~bash
git commit -m "Adiciona novo desafio de lógica"
~~~~
Envie suas modificações para sua branch:
~~~~bash
git push origin nova-feature
~~~~
## Abra um Pull Request.🤖
### 📚 Recursos Úteis
 - Documentação do Golang
 - Documentação do JavaScript
 - Documentação do Delphi
 - Guia SQL
   
## 🛠️ Tecnologias Utilizadas
Este projeto utiliza as seguintes tecnologias:

 - Golang: Versão 1.17+
 - JavaScript: ECMAScript 6+
 - Delphi: Versão XE ou superior
 - SQL: Compatível com MySQL, PostgreSQ, SQLite e SQL Server

## 📧 Contato
Caso tenha dúvidas ou sugestões, entre em contato comigo:

 - Email: gaitolini@gmail.com
 - LinkedIn: [Anderson Gaitolini](https://www.linkedin.com/in/andersongaitolini/)
 - Whatsapp: [Me adicione no WhatsApp](https://wa.me/qr/CFND4RGOJHHUN1)
