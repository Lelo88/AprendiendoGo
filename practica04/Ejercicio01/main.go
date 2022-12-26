package main

import "fmt"

type Product struct {
	Id int64
	Name string
	Price float64
	Description string
	Category string
}

var products =[] Product {
	{
	   Id: 987,
	   Name: "Manzana",
	   Price: 45.6,
	   Description: "",
	   Category: "F",
   },
   {
	   Id: 345,
	   Name: "Carne",
	   Price: 100.53,
	   Description: "Carne de vaca",
	   Category: "C",
   },
}

func (p *Product) Save(product Product) {
	
}

func main(){
	fmt.Println(products)
}