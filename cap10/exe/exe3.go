package exe

import "fmt"

type veiculo struct {
	portas int
	cor    string
}

type caminhonete struct {
	veiculo         veiculo
	tracaoNasQuatro bool
}

type sedan struct {
	veiculo    veiculo
	modeloLuxo bool
}

func MostraCarros() {
	vectra := sedan{veiculo{4, "azul"}, true}

	montana := caminhonete{veiculo{2, "preto"}, true}

	fmt.Println(vectra)
	fmt.Println(montana)
}

/**
- Crie um novo tipo: veiculo
 - O tipo subjacebte deve ser struct
 - Deve conter os campos: porta, cor
- Crie dois novos tipos: caminhonete e sedan
	- Os tipos subjacentes devem ser struct
	- Ambos devem conter "veiculo" como struct embutido
	- O tipo caminhonete deve conter um campo bool chamado "tracaoNasQuatro"
	- O tipo sedan deve conter um campo bool chamado "modeloLuxo"
- Usando os structs veiculo, caminhonete e sedan:
	- Usando composite literal, crie um valor de tipo caminhonete e dê os valores de seus campos
	- Usando composite literal, crie um valor de tipo sedan e dê os valores de seus campos
	- Demonstre estes valores
	- Demonstre um único campo de cada um dos dois
**/
