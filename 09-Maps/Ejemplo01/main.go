package main

import "fmt"

func main(){
	capitales:=map[string]string{
		"España":"Madrid",
		"Argentina":"Buenos Aires",
		"Japon":"Tokio",
	}

	for pais, capital := range capitales{
		fmt.Println("La capital de",pais,"es",capital)
	}
}