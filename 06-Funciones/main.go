package main

import "fmt"

func main(){
	fmt.Println(uno(5))

	numero, estado := dos(2) //3. esta es la manera en la que paso dos argumentos para los parametros que me va a devolver la funcion dos
	fmt.Println(numero)
	fmt.Println(estado)
	fmt.Println("-----------------------------------")
	fmt.Println(Calculo(1,2,3,4,5,6))
	fmt.Println(Calculo(10,20,30))
}

func uno(numero int) int { //1. las funciones reciben parametro y la anotacion fuera del parentesis representa el retorno de lo que vamos a devolver
	return numero *2
}

func dos(numero int)(int, bool){ //2. si devuelvo una lista de tipo de datos, tengo qque volver a ponerlo en parentesis
	if numero ==1{
		return 5, false
	}else{
		return 10,true
	}
}

func Calculo(numero...int) int{ //3. esta es una forma de pasar muchos parametros para devolvernos uno o mas resultados
	total := 0
	for _, num := range numero{
		total = total + num
	}
	return total
}
