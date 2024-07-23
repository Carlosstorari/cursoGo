package exe

import "fmt"

func StructAnonimoExe() {

	x := struct {
		nomes          []string
		sobrenomesECpf map[string]string
	}{
		nomes: []string{"Carlos", "JP"},
		sobrenomesECpf: map[string]string{
			"Strorari": "35453535",
			"Silva":    "6756757",
		},
	}

	fmt.Println(x)
}

/**
- Crie e use um struct anonimo
- Desafio: dentro do struct tenha um
valor de tipo map e outro do tipo slice
**/
