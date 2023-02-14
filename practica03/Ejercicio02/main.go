package main

import (
	"fmt"
)

func main(){
	
	resultado := promedio(6,7,8,9,8,5.6,10)
	
	if resultado <=0 {
		fmt.Println("El promedio no es correcto")
	} else {
		fmt.Printf("%.2f\n", resultado)	
	}
}

func promedio(values...float64)float64{
	resultado := 0.0
	for _,value := range values{
		resultado+=value
		if value < 0{
			return 0
		}
	}
	return resultado/float64(len(values))
}

