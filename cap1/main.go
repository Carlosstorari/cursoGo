package main

import (
	"fmt"

	"github.com/carlosStorari/cursoGo/cap1/exenivel1"
)

type hotdog int

var b hotdog = 101

func main() {
	//fmtPrints()
	//usingTypeCreated()
	//conversionTypeVar()
	//exenivel1.Exe1()
	exenivel1.Exe4()
}

func fmtPrints() {
	x := "oi bom dia\ncomo vai?\tespero \"que\" tudo bem"
	fmt.Println(x)

	oi := "oi"
	bd := "bom dia"
	y := fmt.Sprint(oi, " ", bd)
	fmt.Println(y)
}

func usingTypeCreated() {
	fmt.Printf("%T , %v", b, b)
}

func conversionTypeVar() {
	x := 10
	fmt.Printf("b tipo = %T, valor = %v\n", b, b)
	fmt.Printf("x tipo = %T, valor = %v\n", x, x)
	x = int(b)
	fmt.Printf("b tipo %T <-usando conversion\n", int(b))
	fmt.Printf("x valor = %v  <-novo valor atribuido usando conversion\n", x)

}
