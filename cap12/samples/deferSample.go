package samples

import "fmt"

func SampleDefer() {
	defer fmt.Println("1")
	defer fmt.Println("2")
	fmt.Println("3")
	fmt.Println("4")
}
