package main

import (
	"fmt"
)

var x int = 42
var y string = "James Bond"
var z bool = true

func main() {
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)

	s := fmt.Sprintf("%v\t%v\t %v", x, y, z)
	fmt.Println(s)
}

/****
### EXERCÍCIO 2
- Use var para declarar três variáveis. Elas devem ter package-level scope. Não atribua valores a estas
variáveis. Utilize os seguintes identificadores e tipos para estas variáveis:
	1. Identificador "x" deverá ter tipo int
	2. Identificador "y" deverá ter tipo string
	3. Identificador "z" deverá ter tipo bool
- Na fução main:
	1. Demonstre os valores de cada identificador
	2. O compilador atribuiu valores a essas variaveis.
	Como esses valores se chamam ? R: são chamados de valores 0

### EXERCÍCIO 3
- Utilizando a solução do exercício anterior:
	1. Em package-level scope, atribua os seguintes valores às variáveis:
		1. para "x" atribua 42
		2. para "y" atribua "James Bond"
		3. para "z" atribua true
	2. Na função main:
		1. fmt.Sprintf para atribuir todos esses valores a uma única variável. Faça essa atribuição de tipo
		string a uma variável de nome "s" utilizando o operador curto de declaração
		2. demonstre a variável "s".
****/
