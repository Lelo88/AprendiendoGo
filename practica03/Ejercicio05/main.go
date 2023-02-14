package main

import "fmt"

const (
	dog = "dog"
	cat = "cat"
	hamster = "hamster"
	spider = "spider"
)

type Operation func(cant int) (float64)

func alimentoPerro(perros int) (float64){
	return float64(perros*10)
}

func alimentoGato(gatos int) float64{
	return float64(gatos*5)
}

func alimentoHamster(hamster int) float64{
	return float64(hamster) * 0.25
}

func alimentosArana(arana int) float64{
	return float64(arana) * 0.15
}

func calculaAlimento(animal string) (resultado Operation){
	switch animal{
		case dog: 
			resultado = alimentoPerro
		case cat:
			resultado = alimentoGato
		case hamster:
			resultado = alimentoHamster
		case spider:
			resultado = alimentosArana 
	}
	return 
}

func main(){
	operacion := calculaAlimento("dog")
	fmt.Println(operacion(10))
}