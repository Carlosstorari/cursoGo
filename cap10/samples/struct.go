package samples

import "fmt"

type cliente struct {
	nome      string
	sobrenome string
	fumante   bool
}

func SampleStruct() {
	cliente1 := cliente{
		nome:      "Carlos",
		sobrenome: "Storari",
		fumante:   false,
	}

	cliente2 := cliente{"Joca", "Pereira", true}

	fmt.Println(cliente1)
	fmt.Println(cliente2)
}
