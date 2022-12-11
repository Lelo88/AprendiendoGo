package main

import (
	"fmt"
	"bufio"
	"os"
)

func main(){
	fmt.Println("Ingrese una cadena de texto")
	scanner:=bufio.NewScanner(os.Stdin)
	cadena, cadena2:= "",""
	i, palin, palfin := 0,0,0
	if scanner.Scan(){
		cadena=scanner.Text()
		cadena2 =scanner.Text()
	}

	for i=0;i<len(cadena) && (cadena[i]!='\n'||cadena[i]==' ');i++{
		cadena2 = cadena
	}
	
	palfin = i-1

	for palin<palfin && ((cadena[palin]==cadena2[palfin])||cadena[palin]==' ' || cadena2[palfin]==' '){
		palin++
		palfin--
	}

	if palin>=palfin{
		fmt.Println("La frase ingresada es palindroma")
	}else{
		fmt.Println("La frase ingresada no es palindroma")
	}
}