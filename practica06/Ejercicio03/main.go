package main

import (
	"fmt"
	"errors"
)

var errorSalarioMinimo = errors.New("el salario es menor a 10000")

func calculaSalario(sal int) error{
	if sal<10000{
		return errorSalarioMinimo
	}
	return nil
}

func main(){
	salary := 0
	
	fmt.Println("Ingrese un salario: ")
	fmt.Scanf("%d", &salary)

	err:=calculaSalario(salary)

	if errors.Is(err,errorSalarioMinimo){
		fmt.Println(err)
	}
}