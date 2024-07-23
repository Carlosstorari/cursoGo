package exe

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	sabores   []string
}

func MostraPessoa() {
	pessoa1 := pessoa{
		nome:      "Carlos",
		sobrenome: "Storari",
		sabores:   []string{"morango", "chocolate", "maracuja"},
	}

	fmt.Println(pessoa1)
}

func MostraPessoasNoMap() {

	pessoas := map[string]pessoa{
		"Storari": {
			nome:      "Carlos",
			sobrenome: "Storari",
			sabores:   []string{"morango", "chocolate", "maracuja"},
		},
		"Lopes": {
			nome:      "Vanessa",
			sobrenome: "Lopes",
			sabores:   []string{"pistache", "milho", "hortelã"},
		},
	}

	for _, value := range pessoas {
		fmt.Println("Meu nome é", value.nome, " e meus sorvetes favoritos são: ")
		for _, value := range value.sabores {
			fmt.Println("-", value)
		}
	}
}

/**
### exe1
- Crie um tipo "pessoa" com tipo subjacente struct,
que possa conter os seguintes campos:
	- Nome
	- Sobrenome
	- Sabores favoritos de sorvete
- Crie dois valores do tipo "pessoa" e demonstre
estes valores, utilizando range na slice que contem
os sabores de sorvete.

### exe2
- Utilizando a solução anterior, coloque os valores do tipo
"pessoa" num map, utilizando os sobrenomes como key.
- Demonstre os valores do map utilizando range.
- Os diferentes sabores devem ser demonstrados
utilizando outro range, dentro do range anterior
**/
