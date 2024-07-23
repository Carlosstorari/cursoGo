package samples

import "fmt"

func SampleFuncaoAnonima() {
	x := 387

	func(x int) {
		fmt.Println(x, "vezes 87344 é:")
		fmt.Println(x * 87344)
	}(x)
}
