package main

import "fmt"

func main(){
	paises := make(map[string]string)
	fmt.Println(paises)
	paises["Mexico"] = "D.F."
	paises["Argentina"] = "Buenos Aires"

	fmt.Println(paises["Mexico"]) // comportamiento similar a los diccionarios en python

	fmt.Println(paises)

	//otras formas de crear un map

	//directo con map: map[campo]valor del campo

	campeonato := map[string]int{"Barcelona":39, "Real Madrid":38, "Atletico de Madrid": 37} // asignamos directo

	fmt.Println(campeonato)//si lo imprimo me muestra que se imprime por orden de claves

	campeonato["Getafe"] = 30 // agrego un elemento al mapa
	fmt.Println(campeonato)

	campeonato["Barcelona"] = 40 // modifico el valor de un elemento del mapa
	fmt.Println(campeonato)

	for equipo, puntaje :=  range campeonato{
		fmt.Printf("El equipo %s tiene un puntaje de %d \n", equipo, puntaje)
	}

	//podemos comprobar si un elemento se encuentra en el mapa
	puntaje, existe := campeonato["Independiente"]
	fmt.Printf("El puntaje capturado es %d y el equipo existe %t ", puntaje, existe)
}