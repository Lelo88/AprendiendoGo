package main

import "fmt"

func main(){
	//solo tenemos el ciclo for como iterador, pero hay que agregar la suma del iterado debajo de todo
	/*var i = 0

	for i < 10 {
		fmt.Printf("Valor de i: %d", i)
		if i==5{
			fmt.Println("Multiplicamos por 2")
			i = i*2
			continue
		}
		i++
	}*/

	//vamos a hacer otro ciclo for

	var i int = 0

	RUTINA:
		for i<10{
			if i == 4{
				i = i + 2
				fmt.Println("voy a ir a Rutina")
				goto RUTINA
			}
			fmt.Printf("Valor de i: %d\n", i)
			i++
		}
}