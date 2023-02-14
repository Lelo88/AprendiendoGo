package main

import "testing"

func TestPromedio(t *testing.T){
	var notes = []float64{4,6,5,9,6}
	promedioEsperado := (4.0 + 6.0 + 5.0 + 9.0 + 6.0)/float64(len(notes))

	resultado:=promedio(notes...)

	if promedioEsperado != resultado{
		t.Errorf("La funcion promedio() da como resultado %v. Se esperaba %v", resultado, promedioEsperado)
	} else{
		println("Los promedios coinciden")
	}
}