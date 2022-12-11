package main

import (
	"fmt"
	"strconv"
)

//definicion de variables en go
var numero int
var texto string
var booleano bool = true

func main(){
	fmt.Println("Variables en Go")
	//otro tipo de definicion de variable
	numero2:=2
	
	numero3,numero4, numero5:= 4,5,6 // puedo asignar variables de esta manera

	numero = numero3 + numero4 + numero5 

	fmt.Print(numero2)
	fmt.Println("\n",numero)

	mostrarStatus()

	var moneda float32 = 0
	numero2 = int(moneda)
	texto = strconv.Itoa(int(moneda)) //hay funciones para convertir variables, aunque muchas veces las conversiones no se puedan realizar
	fmt.Println(texto)
}

func mostrarStatus(){
	fmt.Println(booleano)
}
