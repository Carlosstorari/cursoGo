package samples

import "fmt"

func SampleStructAnonimo() {
	x := struct {
		nome  string
		idade int
	}{
		nome:  "Mario",
		idade: 50,
	}

	fmt.Println(x)
}
