package samples

import "fmt"

func SampleMap() {
	//cria map
	amigos := map[string]int{
		"alfredo": 545353,
		"joana":   535451,
	}

	fmt.Println(amigos)
	fmt.Println(amigos["joana"])

	//adiciona novo elemento ao map
	amigos["gopher"] = 4444444

	fmt.Println(amigos)
}

func SampleCommaOkInMap() {
	amigos := map[string]int{
		"alfredo": 545353,
		"joana":   535451,
	}

	//o "ok" depois davirgula verifica se existe elemento "fantasma"
	if verifica, ok := amigos["fantasma"]; !ok {
		fmt.Println("não tem!")
	} else {
		fmt.Println(verifica)
	}
}

func SampleDeleteMapElement() {
	amigos := map[string]int{
		"alfredo": 545353,
		"joana":   535451,
	}

	delete(amigos, "alfredo")

	fmt.Println(amigos)
}

func SampleRangeMap() {
	amigos := map[string]int{
		"alfredo": 545353,
		"joana":   535451,
	}

	//usa o range para percorrer o map
	for key, value := range amigos {
		fmt.Println(key, value)
	}
}
