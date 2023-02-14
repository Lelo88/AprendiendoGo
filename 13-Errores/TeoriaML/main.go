//los errores en go son un tipo de variable interface. Al implementarse, cualquier tipo de dato puede interpretarse como un 
//error. Para esto debemos implementar un metodo error que nos tiene que devolver un string

package main

import "fmt"

type error interface{
	Error() string
}

type MyError struct{
	numError int 
	
}
//implementacion del metodo error

func (e *MyError) Error() string {
	return "my error info"
}

func main(){
	

	fmt.Println()
}