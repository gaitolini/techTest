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

