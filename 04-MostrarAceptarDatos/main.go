package main

import (
	"bufio"
	"fmt"
	"os"
)

var numero1 int
var numero2 int
var leyenda string

func main(){
	
	//el ingreso no funciona si lo ingreso de manera normal, sino como puntero en C
	fmt.Println("Ingrese el número 1:")
	fmt.Scanln(&numero1)
	fmt.Println("Ingr6ese el número 2:")
	fmt.Scanln(&numero2)

	

	scanner:=bufio.NewScanner(os.Stdin)
	if scanner.Scan(){
		leyenda = scanner.Text()
	}
	fmt.Println(numero1 + numero2)
	fmt.Println(leyenda)

}