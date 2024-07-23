package exenivel1

import (
	"fmt"
)

type falafel int

var almoco falafel = 42
var janta int

func Exe4() {
	fmt.Println(almoco)
	janta = int(almoco)
	fmt.Println(janta)
}

/****
 - Crie um tipo. O tipo subjacente deve ser int.
 - Crie uma variável para este tipo, com o identificador "x", utilizando a palavra-chave var.
 - Na função main:
	1. Demonstre o valor da variável "x"
	2. Demonstre o tipo da variável "x"
	3. Atribua 42 à variável "x" utilizando o operator "="
	4. Demonstre o valor da variável "x"
****/
