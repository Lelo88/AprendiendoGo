package main

import (
	//"bufio"
	"fmt"
	//"os"
)

func main(){
/*	var letras string

	fmt.Println("Ingrese una palabra:")
	palabra :=bufio.NewScanner(os.Stdin)
	if palabra.Scan(){
		letras=palabra.Text()
	}

	fmt.Println("La cantidad de letras que tiene la palabra ingresada es:",len(letras))

	for i:=0;i<len(letras);i++{
		fmt.Println(string(letras[i]))
	}
*/

var palabra string
fmt.Println("Ingrese una palabra: ")
fmt.Scanf("%s",&palabra)

fmt.Println("La cantidad de caracteres de la cadena ingresada es de", len(palabra))

	for _,letra:=range palabra{
		fmt.Println(string(letra))
	}
}