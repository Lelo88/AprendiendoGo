package main

import "testing"

func TestCalculaSalario(t *testing.T){
	minutosTrabajados := 456
	categoria := "C"

	salarioEsperado := float64((minutosTrabajados*1000)/60)
	salario:=calculaSalario(minutosTrabajados, categoria)

	if salario != salarioEsperado{
		t.Errorf("La funcion calculaSalario() da como resultado %.2f, pero el resultado esperado es %.2f", salario, salarioEsperado)
	} else{
		println("Salario esperado: %.2f", salarioEsperado, "Salario: %.2f", salario )
	}
}