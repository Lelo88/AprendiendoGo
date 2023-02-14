package main

import "testing"

func TestCalculaImpuesto(t *testing.T){
	sueldo1 ,sueldo2, sueldo3 := 45000.0, 51000.0, 155000.0

	resultadoEsperado1, resultadoEsperado2, resultadoEsperado3 := sueldo1*0.0, sueldo2*0.17, sueldo3*0.27
	
 	resultado1, resultado2, resultado3 :=calculaImpuesto(sueldo1), calculaImpuesto(sueldo2), calculaImpuesto(sueldo3)

	if resultado1 != resultadoEsperado1 && resultado2!=resultadoEsperado2 && resultado3!=resultadoEsperado3{
		t.Errorf("La funcion calcular impuestos no funciona")
	} else {
		println("Resultado Esperado:", resultadoEsperado1," Resultado: ", resultado1)
		println("Resultado Esperado:", resultadoEsperado2," Resultado: ", resultado2)
		println("Resultado Esperado:", resultadoEsperado3," Resultado: ", resultado3)
	}
}
