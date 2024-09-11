# 🚀 Projeto Delphi: Cálculo de Cores de Resistores
Este projeto em Delphi calcula o código de cores de resistores com base em uma string de cor fornecida e inclui uma suite de testes para validar a implementação.

## 🛠️ Requisitos para rodar o projeto
Para rodar este projeto localmente, você precisará:

- Delphi XE ou superior para compilar e rodar o projeto.
- Os arquivos que compõem o projeto, incluindo os testes, já estão anexados.
## 🔗 Onde obter o Delphi?
- Você pode obter o Delphi e suas versões de avaliação no site oficial da Embarcadero:
[Download Delphi](https://www.embarcadero.com/br/products/delphi/starter/free-download)

## 📋 Passo a Passo para rodar o projeto Delphi
### 1️⃣ Instale o Delphi (versão XE ou superior)
- Faça o download e instalação da versão do Delphi no seu computador.
### 2️⃣ Abra o projeto ResistorColor.dproj no Delphi
- Navegue até o arquivo ResistorColor.dproj e abra-o no Delphi. Este arquivo é o projeto principal.
### 3️⃣ Verifique os arquivos do projeto
- Certifique-se de que os seguintes arquivos estão incluídos no projeto:
- uResistorColor.pas: Este arquivo contém a classe TResistor, responsável por calcular o código de cores de resistores.
- uResistorColorTests.pas: Este arquivo contém os testes para verificar a funcionalidade da classe TResistor.
### 4️⃣ Revise o código principal em uResistorColor.pas

Aqui está o código principal:

~~~~delphi
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

### 5️⃣ Verifique os testes em uResistorColorTests.pas
Aqui está o código de teste:

~~~~delphi
unit uResistorColorTests;

interface

uses
  DUnitX.TestFramework, uResistorColor;

type

  [TestFixture]
  TTestResistor = class
  public
    [Test]
    [TestCase('Black', 'black,0')]
    [TestCase('Brown', 'brown,1')]
    [TestCase('Red', 'red,2')]
    [TestCase('Orange', 'orange,3')]
    [TestCase('Yellow', 'yellow,4')]
    [TestCase('Green', 'green,5')]
    [TestCase('Blue', 'blue,6')]
    [TestCase('Violet', 'violet,7')]
    [TestCase('Grey', 'grey,8')]
    [TestCase('White', 'white,9')]
    procedure TestColorCode(const color: string; expected: Integer);

    [Test]
    procedure TestInvalidColor;
  end;

implementation

procedure TTestResistor.TestColorCode(const color: string; expected: Integer);
begin
  Assert.AreEqual(expected, TResistor.colorCode(color));
end;

procedure TTestResistor.TestInvalidColor;
begin
  Assert.WillRaise(
    procedure
    begin
      TResistor.colorCode('invalid');
    end,
    Exception,
    'Cor "invalid" não encontrada no código de cores.');
end;

initialization
  TDUnitX.RegisterTestFixture(TTestResistor);
end.
~~~~
### 6️⃣ Execute os testes 🧪
O arquivo uResistorColorTests.pas contém testes de unidade escritos com DUnitX. Para executar os testes, siga os passos:
- Compile o projeto no Delphi.
- Execute os testes diretamente no Delphi utilizando o DUnitX Framework.
### 7️⃣ Compile e execute o projeto 🏃‍♂️
- Após revisar e ajustar, se necessário, clique em Run (F9) no Delphi para compilar e rodar o projeto e seus testes.
### 🔄 O que aprendemos?
- Como utilizar arrays e propriedades de classe em Delphi para armazenar e acessar o código de cores de resistores.
- Como lançar exceções personalizadas quando uma cor inválida é fornecida.
- Como usar DUnitX para criar e executar testes de unidade para validar o comportamento do código.
## 🚀 Executando o projeto localmente

## Caso queira executar o projeto no seu computador:

- Instale o Delphi versão XE ou superior.
- Abra o arquivo ResistorColor.dproj no Delphi.
- Compile e execute o projeto e os testes de unidade.
---
## 📲 Contato
Caso tenha dúvidas ou sugestões, entre em contato comigo:

 - Email: gaitolini@gmail.com
 - LinkedIn: [Anderson Gaitolini](https://www.linkedin.com/in/andersongaitolini/)
 - Whatsapp: [Me adicione no WhatsApp](https://wa.me/qr/CFND4RGOJHHUN1)   


### Os arquivos de projeto incluídos são:

- ResistorColor.dproj: O arquivo de projeto principal.
- uResistorColor.pas: Contém a implementação da classe de cores de resistores.
- uResistorColorTests.pas: Contém os testes de unidade usando DUnitX.
- Certifique-se de que todos os arquivos estejam no mesmo diretório ao abrir e executar o projeto no Delphi.
