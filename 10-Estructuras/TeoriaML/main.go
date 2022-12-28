//Estructuras

package main

import "fmt"

type Persona struct{ //definicion de una estructura en Go
	nombre 		string
	apellido 	string
	edad 		int
	profesion 	string
	Gustos		Preferencias
}

//tambien puedo usar una estructura como tipo de dato

type Preferencias struct {
	Comidas 	string
	Peliculas 	string
	Deportes	string
}

func main(){
	
	//instanciar una estructura //2. para agregar los gustos, tuve que instanciar por dentro la estructura preferencias
	persona1 := Persona{"Leandro", "Villalba", 34, "Software Developer", Preferencias{"Tortilla", "El señor de los anillos", "Futbol"}}

	//otra manera de instanciar:

	persona2 := Persona{
		nombre: "Mara Romina",
		apellido: "Alvarez",
		edad: 36,
		profesion: "Emprendedora",
	}
	
	//podemos declarar una estructura vacia y luego definirla

	var persona3 Persona

	persona3.nombre = "Luna"

	fmt.Println(persona1)
	fmt.Println(persona2)

	//acceso a campos especificos de la estructura instanciada
	fmt.Println(persona1.nombre)
	fmt.Println(persona1.Gustos.Comidas) //accedo a la comida que se encuentra dentro de la estructura preferencias
	fmt.Println(persona2.nombre)
	fmt.Println(persona3.nombre)

}