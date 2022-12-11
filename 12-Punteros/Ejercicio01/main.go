//programa para ingresar una cadena de caracteres y que no cuente los espacios

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	scanner:=bufio.NewScanner(os.Stdin)
	cadena:=""
	if scanner.Scan(){
		cadena=scanner.Text()
	}
	cont:=0
	
	
	for i:=0;i<len(cadena);i++{
		cont++
		if cadena[i] == ' '{
			cont-=1
		}
	}

	fmt.Println("La cadena ingresada tiene",cont,"caracteres")
}	