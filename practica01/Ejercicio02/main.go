package main

import "fmt"

func main(){
	var temperatura float32 
	var humedad uint8
	var presion uint16

	temperatura = 30.5
	humedad = 98
	presion = 1024

	fmt.Println("La temperatura es de ",temperatura,"grados, con una humedad de", humedad,"% y con una presion de", presion,"hpa")
}