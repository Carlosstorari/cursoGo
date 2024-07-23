package samples

import "fmt"

type pessoa struct {
	nome  string
	idade int
}

type profissional struct {
	pessoa
	titulo  string
	salario int
}

func SampleStructEmbutido() {
	pessoa1 := pessoa{
		nome:  "Alfedro",
		idade: 30,
	}

	profissional1 := profissional{
		pessoa: pessoa{
			nome:  "Maricota",
			idade: 31,
		},
		titulo:  "Pizzaiola",
		salario: 10000,
	}

	fmt.Println(pessoa1)
	fmt.Println(profissional1)
}
