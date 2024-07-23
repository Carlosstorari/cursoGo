package samples

import "fmt"

func SampleSwitch() {

	switch x := 2; {
	case x == 1:
		fmt.Println("é um 1")
	case x == 2:
		fmt.Println("é um 2")
	case x == 3:
		fmt.Println("é um 3")
	case x == 4:
		fmt.Println("é um 4")
	}
}
