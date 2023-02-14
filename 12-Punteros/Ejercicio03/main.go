package main

import (
	"fmt"
)

func main(){
	numero:= 0.0
	pnumero := &numero

	fmt.Println("Ingrese un numero con coma: ")
	fmt.Scanf("%f",&numero)

	fmt.Println("Ubicacion de la variable pnumero:", pnumero)	
	fmt.Println("Valor de la variable numero:",*pnumero)

}