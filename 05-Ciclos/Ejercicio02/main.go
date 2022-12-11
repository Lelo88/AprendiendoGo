//programa para averiguar si un numero ingresado es capicua

package main

import "fmt"

func main(){
	numero, numero2, ultimo, reverso := 0,0,0,0
	fmt.Println("Ingrese un numero: ")
	fmt.Scanln(&numero)

	numero2 = numero

	for numero>0{ //bucle for similar a un while
		ultimo = numero%10
		reverso = reverso*10 + ultimo
		numero/=10
	}

	if numero2==reverso{
		fmt.Println("El numero ingresado es capicua")
	}else{
		fmt.Println("El numero ingresado no es capicua")
	}
}