package main

import (
	"fmt"
	"errors"
)

func main(){
	fmt.Println()
	error:=errors.New("este es un error")
	fmt.Println(error)
}