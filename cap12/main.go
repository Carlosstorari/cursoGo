package main

import (
	"github.com/carlosStorari/cursoGo/cap12/samples"
)

func main() {

	//total, quantos, oi := samples.SomaMultiplosRetornos("tarde", 4, 6, 70, 14)

	//fmt.Println(total, quantos, oi)

	//passando slice como argumento
	//si := []int{10, 10, 1, 2, 3, 5}
	//total2, quantos2, oi2 := samples.SomaMultiplosRetornos("tarde", si...)

	//fmt.Println(total2, quantos2, oi2)

	//samples.SampleDefer()

	//samples.MetodosDeTipoSample()
	//samples.SampleInterface()
	//samples.SampleFuncaoAnonima()

	//////FUNÇÃO COMO EXPREÇÃO/////
	// y := func(x int) int {
	// 	return x * 87979
	// }
	// fmt.Println(y(3))
	//////////////////////////////
	samples.SampleRetornaFunc()

}
