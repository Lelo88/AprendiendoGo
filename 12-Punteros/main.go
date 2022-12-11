//Teoría de punteros

package main

import "fmt"

func main(){
	i:=10
	p:=&i
	fmt.Println("El valor de i es",i,"y el puntero p que vale",*p,"apunta a",p)
	a:=*p
	*p=11
	fmt.Println("El valor de a es",a,"y el puntero p que apunta a",p,"ahora vale",*p)

	b:=0
	j:=0
	p1:=&b
	p2:=&j

	if p1==p2{
		fmt.Println("Los punteros p1 y p2 apuntan a la misma direccion")
	}else{
		fmt.Println("Los punteros p1 y p2 NO apuntan a la misma direccion")
	}

	if *p1==*p2 {
		fmt.Println("Los punteros p1 y p2 valen lo mismo")
	}else{
		fmt.Println("Los punteros NO valen lo mismo")
	}
}