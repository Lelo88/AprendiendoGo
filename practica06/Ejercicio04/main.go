package main

import "fmt"

func calculaSalario(sal int) error{
	if sal < 150000 {
		return fmt.Errorf("error: el minimo no imponible es de 150000. El salario es de %d", sal)
	}
	return nil
}

func main(){
	salary:=0
	fmt.Println("Ingrese un salario: ")
	fmt.Scanf("%d", &salary)

	err:=calculaSalario(salary)

	if err != nil{
		fmt.Println(err)
	} else {
		fmt.Println("Debe pagar impuestos")
	}
}