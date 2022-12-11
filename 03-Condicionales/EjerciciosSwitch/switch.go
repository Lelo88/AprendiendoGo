package main

import "fmt"

func main(){
	var c int8
	fmt.Scanf("%c",&c)

	switch {
	case c>='A' && c<='Z':
		fmt.Println("Letra Mayuscula")
	case c>='a' && c<='z':
		fmt.Println("Letra Minuscula")
	case c>='0' && c<='9':
		fmt.Println("Número")
	default:
		fmt.Println("Otro caracter")
	}

}