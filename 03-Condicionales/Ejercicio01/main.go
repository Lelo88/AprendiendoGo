package main

import "fmt"

func main(){
	opcion, num1, num2:=0,0,0
	otrosNum:=false
	for{
		if opcion!=5{
			for !otrosNum{
				fmt.Println("Ingrese el primer número: ")
				fmt.Scanln(&num1)
				fmt.Println("Ingrese el segundo número: ")
				fmt.Scanln(&num2)
				otrosNum = true
			}
			fmt.Println("1- Suma")
			fmt.Println("2- Resta")
			fmt.Println("3- Multiplicacion")
			fmt.Println("4- Division")
			fmt.Print("Ingrese una opción: ")
			fmt.Scanln(&opcion)
			switch opcion{
			case 1: 
			fmt.Println("Suma de los números ingresados: ", num1 + num2)
			otrosNum = false
			case 2:
			fmt.Println("Resta de los números ingresados: ", num1-num2)
			otrosNum = false
			case 3: 
			fmt.Println("Multiplicacion: ", num1*num2)
			otrosNum = false
			case 4:
			fmt.Println("Division: ", num1/num2)
			otrosNum = false
			case 5:
			fmt.Println("Usted esta saliendo del programa")
			default:
			fmt.Println("Opcion incorrecta")
			}
		}else{
			break
		}
	}
}