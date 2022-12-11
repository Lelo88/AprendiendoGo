package main

import "fmt"

var calculo func(int, int) int = func(num1 int, num2 int) int{
	return num1 + num2 //funcion anonima. A diferencia de las funciones normales, puedo modificarla
}

func main(){
	fmt.Printf("Sumo 5 + 7 = %d \n", calculo(5,7))

	calculo = func(num1 int, num2 int) int{ //puedo modificar la funcion anonima
		return num1 - num2
	}

	fmt.Printf("Resto 7-5: %d \n", calculo(7,5))

	calculo = func(num1 int, num2 int) int{ //puedo modificar la funcion anonima
		return num1/num2
	}

	fmt.Printf("Division 7-5: %d \n", calculo(7,5))


	fmt.Println("-----------------")
	Operaciones()
}

func Operaciones(){ //1. Ejemplo de closure sencillo (buscar despues la definicion)
	resultado := func() int{
		var a int = 23
		var b int = 13
		return a + b		
	}
	fmt.Println(resultado())

	tablaDel := 2
	MiTabla := Tabla(tablaDel)
	for i:=1;i<10;i++{
		fmt.Println(MiTabla())
	}
}

func Tabla(valor int) func() int{ // ejemplo de closure mas complicado, donde va a devolver una secuencia. 
	numero := valor					//arriba de todo lo usamos con un for para que haga un recorrido y realice
	secuencia := 0					//lo que se encuentra dentro de este return 
	return func() int {
		secuencia += 1
		return numero * secuencia
	}
}