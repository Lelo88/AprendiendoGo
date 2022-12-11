package main

import "fmt"

var estado bool

func main(){
	estado = true 
	if estado=false; estado == true {
		fmt.Println(estado)
		fmt.Println("Este es el condicional if de go")
	}else{
		fmt.Println(estado)
		fmt.Println("Es es el else de go")
	}

	//switch: no hace falta el break en los cases

	switch numero:=5; numero{
	case 1:
		fmt.Println(numero)
	case 2:
		fmt.Println(numero)
	case 3:
		fmt.Println(numero)
	case 5:
		fmt.Println(numero)
	default:
		fmt.Println("Mayor a 5")
	}
}

