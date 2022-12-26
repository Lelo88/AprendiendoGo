package main

import "fmt"

type Person struct {
	ID 			int
	Name 		string
	DateOfBirth string
}

type Employee struct {
	ID 			int
	Position 	string
	Person
}

func (e *Employee) PrintEmployee(){
	fmt.Printf("ID Persona: %d, ID Empleado: %d, Position: %s, Name: %s, Fecha de nacimiento: %s\n", e.Person.ID, e.ID, e.Position,e.Person.Name, e.Person.DateOfBirth)
}

func main(){
	p:= Person{ID: 345, Name: "Leandro", DateOfBirth: "30/07/1988"}
	e:=	Employee{ID: 12, Position: "Software Developer", Person: p}
	e.PrintEmployee()
}