package main

import "fmt"

func calculaImpuesto(sueldo float64) float64{
	if sueldo>=50000 && sueldo<150000{
		return (sueldo*0.17)
	} else if sueldo>=150000 {
		return (sueldo*0.27)
	} else{
		return 0
	}
}
func main() {
	sueldo:=0.0
	fmt.Println("Ingrese el sueldo del empleado: ")
	fmt.Scanf("%f", &sueldo)

	fmt.Printf("El impuesto a cobrar es de %.2f\n", calculaImpuesto(sueldo))
}

