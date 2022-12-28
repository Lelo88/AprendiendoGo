package main

import (
	"fmt"
	"errors"
)

type ErrorSalario struct{}

func (e ErrorSalario) Error() string {
	return "Error: el salario es menor a 10000"
}

func calculaSalario(salario int) error{
	if salario < 10000 {
		return ErrorSalario{}
	}

	return nil
}

func main(){
	salary := 0
	fmt.Println("Ingrese un salario: ")
	fmt.Scanf("%d", &salary)

	err := calculaSalario(salary)
	
	if errors.Is(err, ErrorSalario{}){
		fmt.Println(err)
	}
}