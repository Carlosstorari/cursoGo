package samples

import "fmt"

func SampleSlice() {
	mySlice := []int{1, 2, 3, 4, 5}

	fmt.Println(mySlice)

	// diferente do array o slice é dinamico e pode almentar de tamanho
	mySlice2 := append(mySlice, 6)

	fmt.Println(mySlice2)
}

func SampleSliceRange() {
	mySlice := []string{"banana", "maçã", "jaca", "pêssego"}

	//range ajuda a percorrer o slice
	for _, valor := range mySlice {
		fmt.Printf("Um dos valores desse slice é %v.\n", valor)
	}
}

func SampleSliceOfSlice() {
	sabores := []string{"pepperoni", "mozzarela", "abacaxi", "quatro queijos", "marguerita"}

	fatia := sabores[3:5]

	fmt.Println(fatia)
}

func SampleDeleteElementOfSlice() {
	sabores := []string{"pepperoni", "mozzarela", "abacaxi", "quatro queijos", "marguerita"}

	//delete o elemento "abacaxi"
	// "..." no final de sabores é para dizer que os elementos em si estao colocados lá
	sabores = append(sabores[:2], sabores[3:]...)

	fmt.Println(sabores)
}

func SampleMakeSlice() {

	//inicializa slice com comprimento 5 e capacidade maxima de 10 elementos
	slice := make([]int, 5, 10)

	slice[0], slice[1], slice[2], slice[3], slice[4] = 1, 2, 3, 4, 5

	slice = append(slice, 6)
	slice = append(slice, 7)
	slice = append(slice, 8)
	slice = append(slice, 9)
	slice = append(slice, 10)

	fmt.Println(slice, len(slice), cap(slice))

	//Ao passar a capacidade maxima o a slice duplica a capacidade pra ser dinamica
	slice = append(slice, 10)

	fmt.Println(slice, len(slice), cap(slice))
}
