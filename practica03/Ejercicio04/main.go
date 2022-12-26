package main

import "fmt"

const (
	minimum = "minimum"
	maximum = "maximum"
	average = "average"
)

type Operation func (notas ...float64) float64

func minimo(notas ...float64) (min float64){
	min = notas[0]
	for _,nota:=range notas{
		if nota<min{
			min = nota
		}
	}
	return min
}

func maximo(notas ...float64) (max float64){
	max = notas[0]
	for _,nota:=range notas{
		if nota>max{
			max = nota
		}
	}
	return max
}

func calculaPromedio(notas...float64) float64{
	suma := 0.0
	for _,nota := range notas{
		suma+=nota
	}
	return suma/float64(len(notas))
}

func operacionProfesores(calculo string) (resultado Operation){
	switch calculo{
	case minimum:
		resultado = minimo
	case maximum: 
		resultado = maximo
	case average:
		resultado = calculaPromedio	
	}

	return resultado
}

func main(){
	operation := operacionProfesores("minimum")
	fmt.Println(operation(2.0,3.9,6.7,7.2,1.1))
}