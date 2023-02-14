package main

import "fmt"

func main(){
	minutos, categoria :=0, ""
	fmt.Println("Ingrese la categoría del empleado: ")
	fmt.Scanf("%s", &categoria)
	fmt.Println("Ingrese la cantidad de minutos trabajados")
	fmt.Scanf("%d", &minutos)

	
	fmt.Printf("El salario correspondiente a los %d minutos trabajados es de %.2f\n", minutos,calculaSalario(minutos, categoria))
}

func calculaSalario(min int, categoria string) (salario float64){
	switch categoria{
	case "C":
		salario = float64((1000 * min)/60)
	case "B":
		salario = float64((1500 * min)/60)
		salario += salario*0.2
	case "A":
		salario = float64((1000 * min)/60)
		salario += salario*0.5
	default:
		fmt.Println("No existe la categoria")
	}
	return salario
}
