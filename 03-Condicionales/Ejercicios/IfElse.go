package main

import (
	"fmt"
	"math/rand"
)

func main(){
	valor:=rand.Int()

	if valor % 2 == 0{
		fmt.Println("El valor",valor,"es par")
	} else{
		fmt.Println("El valor",valor, "no es par")
	}
}