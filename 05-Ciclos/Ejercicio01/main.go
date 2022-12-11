//bucle infinito de salida

package main

import "fmt"

func main(){
	for { // vendria a ser un while infinito. hasta que no rompamos la condicion, se sigue ejecutando 
		fmt.Print("Salir? s/n")
		var c rune
		fmt.Scanf("%c\n", &c)
		if c=='S' || c=='s'{
			break
		}
	}
}