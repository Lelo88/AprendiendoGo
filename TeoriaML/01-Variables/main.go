package main

import "fmt"

func main(){
	//declaración de variables 
	var (
		producto = "carne"
		otroNumero = 15
		verdaderto = true
	)
	var horas int = 0 
	var saludo string = "hola"
	despedida := "adios"
	horas = 15
	minutos := 49 //declaracion corta de una variable
	fmt.Println(horas, minutos)
	fmt.Println(saludo, despedida)
	fmt.Println(producto, otroNumero, verdaderto)
}