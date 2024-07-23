package samples

import "fmt"

type pessoa1 struct {
	nome  string
	idade int
}

func (p pessoa1) oibomdia() {
	fmt.Println(p.nome, "diz bom dia!")
}

func MetodosDeTipoSample() {
	carlos := pessoa1{"Carlos", 30}
	carlos.oibomdia()
}
