package main

import "fmt"

var tabla[10] int
var matriz [5][7] int 

//el vector y matriz definida anteriormente estan definidos de manera estatica. se pueden definir de manera
//dinamica

func main(){
	tabla[0] = 1
	tabla[5] = 15
	fmt.Println(tabla)
	fmt.Println("----------------------------")
	tabla2 := [10]int{9,6,8,45,12,36,98,78,145,21}
	fmt.Println(tabla2)
	fmt.Println("----------------------------")
	for i:=0;i<len(tabla2);i++{
		fmt.Println(tabla2[i])
	}
	fmt.Println("----------------------------")

	matriz[3][5] = 1

	fmt.Println(matriz)

	fmt.Println("--------------------")

	matriz2:=[]int{2,5,4} // slice, porque inicialmente no le he puesto ninguna dimension
	fmt.Println(matriz2)

	fmt.Println("--------------------")

	variante2()

	fmt.Println("--------------------")

	variante3()

	fmt.Println("--------------------")

	variante4()
}

func variante2(){
	elementos :=[5]int{1,2,3,4,5}
	porcion := elementos[2:] //slice que guarda una parte del vector anterior

	fmt.Println(porcion)
}

func variante3(){
	elementos := make([]int,5,20) //crea un vector de 5 pero me reserva un espacio de 20
	fmt.Printf("Largo :%d, Capacidad %d", len(elementos), cap(elementos))
}

func variante4(){
	nums:= make([]int, 0,0)
	for i:=0;i<100;i++{
		nums = append(nums, i) // con append voy agregando mas espacio mientras mas lo necesite, por lo que tiene su costo
	}

	fmt.Printf("\n Largo: %d, Capacidad: %d", len(nums), cap(nums))
}