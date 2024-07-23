package samples

import "fmt"

func SampleRetornaFunc() {
	x := retornaFunc()
	y := x(3)
	fmt.Println(y)
	////////////////////////////////////
	//outra forma de retornar o resultado
	fmt.Println(retornaFunc()(3))
}

func retornaFunc() func(int) int {
	return func(i int) int {
		return i * 10
	}
}
